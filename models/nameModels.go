package models

import "github.com/gofiber/fiber/v2"

type Name struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

var names = []Name{}

func CreateName(c *fiber.Ctx) error {
	newName := new(Name)

	err := c.BodyParser(newName)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	names = append(names, *newName)
	return c.Status(201).JSON(newName)
}

func GetNames(c *fiber.Ctx) error {
	return c.JSON(names)
}
