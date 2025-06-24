package controller

import (
	"api.fiber.practice/models"
	"github.com/gofiber/fiber/v2"
)

var names = []models.FullNameRet{}

func HandlerUpdateName(c *fiber.Ctx) error {
	// Getting the Full Name first
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

	// Parse the New Name
	newName := new(models.FullNameRet)
	err = c.BodyParser(newName)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    err.Error(),
		})
	}

	if newName.FirstName == "" && newName.MiddleName == "" && newName.LastName == "" {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    "Name Required",
		})
	}

	if newName.FirstName != "" {
		names[i-1].FirstName = newName.FirstName
	}

	if newName.MiddleName != "" {
		names[i-1].MiddleName = newName.MiddleName
	}

	if newName.LastName != "" {
		names[i-1].LastName = newName.LastName
	}

	return c.Status(400).JSON(fiber.Map{
		"code":   201,
		"status": "success",
		"msg":    names[i-1],
	})
}

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
