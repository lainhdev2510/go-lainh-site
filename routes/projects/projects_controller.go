package projects

import (
	"go-lainh-site/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterProjects(app *fiber.App) {
	app.Get("/projects/:projectId", func(c *fiber.Ctx) error {
		project := utils.FetchProject(c.Params("projectId"))
		return utils.Render(c, ProjectPage(project))
	})
	app.Get("/projects", func(c *fiber.Ctx) error {
		projects := utils.FetchProjects()
		return utils.Render(c, ProjectsPage(projects))
	})
}
