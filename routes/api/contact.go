package api

import "github.com/gofiber/fiber/v2"

func ContactHandler(c *fiber.Ctx) error {
	return c.SendString("Thanks for contacting me")
}
