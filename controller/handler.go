package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandlerGet(c *fiber.Ctx) error {
	fmt.Print("Practicing Get")

	return nil
}

func HandlerPost(c *fiber.Ctx) error {

	return nil
}
