package home

import (
	"go-lainh-site/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterHome(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		blogs := utils.FetchFeaturedBlogs()
		projects := utils.FetchFeaturedProject()
		return utils.Render(c, HomePage(blogs, projects))
	})
}
