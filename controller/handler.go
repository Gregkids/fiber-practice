package controller

import (
	"api.fiber.practice/models"
	"github.com/gofiber/fiber/v2"
)

func HandlerGetNames(c *fiber.Ctx) error {
	return c.JSON(models.TestGetNames)
}

func HandlerCreateName(c *fiber.Ctx) error {

	return nil
}
