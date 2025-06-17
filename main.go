package main

import (
	"api.fiber.practice/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/get", controller.HandlerGet)
	app.Post("/post", controller.HandlerPost)

	app.Listen(":3000")
}
