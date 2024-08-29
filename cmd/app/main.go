package main

import (
	"go-lainh-site/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	app := fiber.New(fiber.Config{
		Network: fiber.NetworkTCP,
	})

	app.Use(logger.New())
	app.Static("/public", "./public")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Set("Surrogate-Control", "no-store")
		return c.Next()
	})

	routes.RegisterRoutes(app)

	addr := os.Getenv("HTTP_LISTEN_ADDR")

	log.Fatal(app.Listen(addr))
}
