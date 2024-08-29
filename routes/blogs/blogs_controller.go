package blogs

import (
	"go-lainh-site/utils"

	"github.com/gofiber/fiber/v2"
)

func BlogsRegiser(app *fiber.App) {

	app.Get("/blogs/:blogID", func(c *fiber.Ctx) error {
		blog := utils.FetchBlog(c.Params("blogID"))
		return utils.Render(c, BlogPage(blog))
	})
	app.Get("/blogs", func(c *fiber.Ctx) error {
		blogs := utils.FetchBlogs()
		return utils.Render(c, BlogsPage(blogs))
	})
}
