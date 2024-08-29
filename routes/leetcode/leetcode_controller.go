package leetcode

import (
	"go-lainh-site/utils"

	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/yuin/goldmark"
)

func LeetcodeRegister(app *fiber.App) {
	app.Get("/leetcode/:problemId", func(c *fiber.Ctx) error {
		problem := utils.FetchProblem(c.Params("problemId"))
		var buf bytes.Buffer
		if err := goldmark.Convert([]byte(problem.Content), &buf); err != nil {
			panic(err)
		}
		return utils.Render(c, ProblemPage(problem, buf.String()))
	})

	app.Get("/leetcode", func(c *fiber.Ctx) error {
		problems := utils.FetchAllProblem()
		return utils.Render(c, LeetcodePage(problems))
	})
}
