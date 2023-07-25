package controller

import "github.com/gofiber/fiber/v2"

func InitialController() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jsonData := fiber.Map{"Name": "Pub.dev Package Server", "Created By": "Koustav Mondal <ThunderBlast>", "Status": "Running", "Version": "1.0.0"}
		return ctx.Status(fiber.StatusOK).JSON(jsonData)
	}
}
