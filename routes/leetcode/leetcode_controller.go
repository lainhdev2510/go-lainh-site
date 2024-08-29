package leetcode

import (
	"go-lainh-site/utils"

	"github.com/gofiber/fiber/v2"
)

func LeetcodeRegister(app *fiber.App) {
	app.Get("/leetcode/:problemId", func(c *fiber.Ctx) error {
		problem := utils.FetchProblem(c.Params("problemId"))
		return utils.Render(c, ProblemPage(problem))
	})

	app.Get("/leetcode", func(c *fiber.Ctx) error {
		problems := utils.FetchAllProblem()
		return utils.Render(c, LeetcodePage(problems))
	})
}
