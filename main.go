package main

import (
	"api.fiber.practice/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/names", controller.HandlerGetNames)
	app.Get("/names/:id", controller.HandlerGetOneName)
	app.Post("/add-name", controller.HandlerCreateName)

	app.Listen(":3000")
}
