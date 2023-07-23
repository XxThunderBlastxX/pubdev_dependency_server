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
		page := ctx.Query("page", "1")

		// Replace all the spaces with +. This is required for the url as if it contains spaces it won't work
		q = strings.ReplaceAll(q, " ", "+")

		// Create a new collector
		c := colly.NewCollector(colly.AllowedDomains("pub.dev"))

		// On every HTML element with class packages-item execute the following
		c.OnHTML(".packages-item", func(e *colly.HTMLElement) {
			// Get the DOM of the HTML
			d := e.DOM

			// Get the title of the package
			title := d.Find(".packages-title").Text()

			// Get the description of the package
			desc := d.Find(".packages-description").Text()

			// Get the metadata of the package
			metadata := d.Find("p.packages-metadata")

			// Get the version of the package from metadata
			version := metadata.ChildrenFiltered("span:contains(ago)").Text()

			// Get the number of likes
			likes, _ := strconv.Atoi(d.Find(".packages-score-like").Children().Find("span").Text())

			// Get the number of pub points
			pubPoints, _ := strconv.Atoi(d.Find(".packages-score-health").Children().Find("span").Text())

			// Get the popularity of the package
			popularityAsString := d.Find(".packages-score-popularity").Children().Find("span").Text()
			popularityAsString = strings.ReplaceAll(popularityAsString, "%", "")
			popularity, _ := strconv.Atoi(popularityAsString)

			// Get the image url of the package
			imgUrl := d.Find(".thumbnail-image").AttrOr("src", "")

			// Create a new package
			pkg := Package{
				Title:      title,
				Desc:       desc,
				Version:    version,
				Likes:      likes,
				PubPoints:  pubPoints,
				Popularity: popularity,
				ImgUrl:     imgUrl,
			}

			// Append the package to the packages list
			packages.Packages = append(packages.Packages, pkg)
		})

		// Visit the url to search the query
		c.Visit(fmt.Sprintf("https://pub.dev/packages?q=%s&page=%s", q, page))

		// Return the packages as JSON
		return ctx.Status(fiber.StatusOK).JSON(packages)
	}
}
