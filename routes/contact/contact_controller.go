package contact

import (
	"go-lainh-site/utils"

	"github.com/gofiber/fiber/v2"
)

func ContactRegister(app *fiber.App) {
	app.Get("/contact", func(c *fiber.Ctx) error {
		return utils.Render(c, ContactPage())
	})
}
