package main

import (
	"api.fiber.practice/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/names", controller.HandlerGetNames)

	app.Listen(":3000")
}
