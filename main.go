package main

import (
	"XxThunderBlastxX/pubdev_dependency_server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Loading env variables
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Fiber app instance
	app := fiber.New()

	app.Use(cors.New())

	// Route of the application
	router.Routes(app)

	// Listen to the port
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
