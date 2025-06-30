package controller

import (
	"api.fiber.practice/models"
	"github.com/gofiber/fiber/v2"
)

var names = []models.FullNameRet{}

func HandlerDeleteName(c *fiber.Ctx) error {
	i, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    err.Error(),
		})
	}

	if i > len(names) {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    "No Name Found",
		})
	}

	names = append(names[:i-1], names[i:]...)

	return c.Status(400).JSON(fiber.Map{
		"code":   201,
		"status": "success",
		"msg":    "Name Removed",
	})
}
