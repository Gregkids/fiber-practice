package main

import (
	"api.fiber.practice/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/names", controller.HandlerGetNames)
	app.Get("/name/:id", controller.HandlerGetOneName)
	app.Post("/add-name", controller.HandlerCreateName)
	app.Put("/change-name/:id", controller.HandlerUpdateName)
	app.Delete("/delete-name/:id", controller.HandlerDeleteName)

	app.Listen(":3000")
}
