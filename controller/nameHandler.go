package controller

import (
	"api.fiber.practice/models"
	"github.com/gofiber/fiber/v2"
)

var names = []models.FullName{}

func HandlerGetNames(c *fiber.Ctx) error {
	return c.JSON(names)
}

func HandlerCreateName(c *fiber.Ctx) error {
	newName := new(models.FullName)

	err := c.BodyParser(newName)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    err.Error(),
		})
	}

	if newName.FirstName == "" && newName.LastName == "" {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    "Data Required",
		})
	}

	names = append(names, *newName)

	return c.Status(400).JSON(fiber.Map{
		"code":   201,
		"status": "success",
		"msg":    "Name Added",
	})
}
