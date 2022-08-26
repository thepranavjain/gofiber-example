package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", handler)

	app.Listen(":3000")
}

func handler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}
