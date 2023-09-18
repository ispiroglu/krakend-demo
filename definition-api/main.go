package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	app.Post("/definition", func(c *fiber.Ctx) error {

		log.Info(string(c.Body()))
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
