package router

import (
	"XxThunderBlastxX/pubdev_dependency_server/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controller.InitialController())

	app.Get("/search", controller.PubdevPackageController())
}
