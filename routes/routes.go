package routes

import (
	"go-lainh-site/routes/api"
	"go-lainh-site/routes/blogs"
	"go-lainh-site/routes/contact"
	"go-lainh-site/routes/home"
	"go-lainh-site/routes/leetcode"
	"go-lainh-site/routes/projects"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	home.RegisterHome(app)
	projects.RegisterProjects(app)
	blogs.BlogsRegiser(app)
	leetcode.LeetcodeRegister(app)
	contact.ContactRegister(app)

	api.RegisterApiRoutes(app)
}
