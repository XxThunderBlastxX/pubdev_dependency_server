package controller

import (
	"XxThunderBlastxX/pubdev_dependency_server/models"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// PackageController function to get packages from pub.dev for a given query
func PackageController() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// packages contains the list of packages
		var packages models.PackageList

		// Receive the query from the request
		q := ctx.Query("q")
		page := ctx.Query("page", "1")

		// Replace all the spaces with +. This is required for the url as if it contains spaces it won't work
		q = strings.ReplaceAll(q, " ", "+")

		packages = models.CrawlPackage(q, page)

		// Return the packages as JSON
		return ctx.Status(fiber.StatusOK).JSON(packages)
	}
}
