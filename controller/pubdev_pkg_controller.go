package controller

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

// Package struct
type Package struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Version    string `json:"version"`
	Likes      int    `json:"likes"`
	PubPoints  int    `json:"pubPoints"`
	Popularity int    `json:"popularity"`
	ImgUrl     string `json:"imgUrl"`
}

// PackageList struct
type PackageList struct {
	Packages []Package `json:"packages"`
}

// PubdevPackageController function to get packages from pub.dev for a given query
func PubdevPackageController() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// packages contains the list of packages
		var packages PackageList

		// Receive the query from the request
		q := ctx.Query("q")

		// Create a new collector
		c := colly.NewCollector(colly.AllowedDomains("pub.dev"))

		// On every HTML element with class packages-item execute the following
		c.OnHTML(".packages-item", func(e *colly.HTMLElement) {
			d := e.DOM

			title := d.Find(".packages-title").Text()
			desc := d.Find(".packages-description").Text()
			version := d.Find("p.packages-metadata").ChildrenFiltered("span:contains(ago)").Text()
			likes, _ := strconv.Atoi(d.Find(".packages-score-like").Children().Find("span").Text())
			pubPoints, _ := strconv.Atoi(d.Find(".packages-score-health").Children().Find("span").Text())
			popularity, _ := strconv.Atoi(strings.Replace(d.Find(".packages-score-popularity").Children().Find("span").Text(), "%", "", -1))
			imgUrl := d.Find(".thumbnail-image").AttrOr("src", "")

			pkg := Package{
				Title:      title,
				Desc:       desc,
				Version:    version,
				Likes:      likes,
				PubPoints:  pubPoints,
				Popularity: popularity,
				ImgUrl:     imgUrl,
			}

			packages.Packages = append(packages.Packages, pkg)
		})

		// Visit the url to search the query
		c.Visit(fmt.Sprintf("https://pub.dev/packages?q=%s", q))

		// Return the packages as JSON
		return ctx.Status(fiber.StatusOK).JSON(packages)
	}
}
