package api

import "github.com/gofiber/fiber/v2"

func RegisterApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/contact", func(c *fiber.Ctx) error {
		return ContactHandler(c)
	})
}
