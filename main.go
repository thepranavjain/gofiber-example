package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", handler)

	app.Get("/with-param/:name?", withParamHandler)

	app.Get("/wildcard/*", wildCardHandler)

	app.Listen(":3000")
}

func handler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func withParamHandler(c *fiber.Ctx) error {
	if c.Params("name") != "" {
		return c.SendString("Hello " + c.Params("name"))
	}
	return c.SendString("Who are you?")
}

func wildCardHandler(c *fiber.Ctx) error {
	return c.SendString("API path: " + c.Params("*"))
}
