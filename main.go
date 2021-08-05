package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		// msg := fmt.Sprintf("x %s", c.Params("*"))
		return c.SendString("Hello")
	})
	app.Listen(":3000")
}
