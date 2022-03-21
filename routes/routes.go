package routes

import (
	"github_api/apps"

	"github.com/gofiber/fiber/v2"
)

func first(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("OK!")
}

func Data(app *fiber.App) {
	app.Get("/", first)
	app.Get("/api/yoga3911/indogram", apps.Indogram)
	app.Get("/api/yoga3911/breadify", apps.Breadify)
}