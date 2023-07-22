package controller

import "github.com/gofiber/fiber/v2"

func PubdevPackageController() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jsonData := fiber.Map{"Name": "Pub.dev Package Server", "Created By": "Koustav Mondal <ThunderBlast>", "Status": "Running", "Version": "0.0.1"}

		return ctx.Status(200).JSON(jsonData)
	}
}
