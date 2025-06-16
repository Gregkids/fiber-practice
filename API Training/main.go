package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/get/:name/:age", nil)
	app.Post("/post/:name/:age", nil)

	app.Listen(":3000")
}
