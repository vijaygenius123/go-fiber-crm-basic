package main

import (
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(8000)
}
