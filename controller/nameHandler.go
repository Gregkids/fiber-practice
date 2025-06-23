package controller

import (
	"database/sql"

	"api.fiber.practice/models"
	"api.fiber.practice/repository"
	"github.com/gofiber/fiber/v2"
)

var names = []models.FullNameReq{
	{FirstName: "Gian", MiddleName: "Indra", LastName: "Nugraha"},
	{FirstName: "John", LastName: "Doe"},
}

func HandlerCreateName(c *fiber.Ctx) error {
	newName := new(models.FullNameReq)

	err := c.BodyParser(newName)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    err.Error(),
		})
	}

	if newName.FirstName == "" {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    "Name Required",
		})
	}

	names = append(names, *newName)

	return c.Status(400).JSON(fiber.Map{
		"code":   201,
		"status": "success",
		"msg":    "Name Added",
	})
}

func HandlerGetNames(c *fiber.Ctx) error {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=darageta dbname=gofiber_test sslmode=disable")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":   500,
			"status": "SQL Error",
			"msg":    err.Error(),
		})
	}

	// Parse Database
	r := repository.NameSQL{DB: db}

	ret, err := r.DBGetAllName()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "error",
			"msg":    err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":   200,
		"status": "success",
		"msg":    ret,
	})
}

func HandlerGetOneName(c *fiber.Ctx) error {
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

	return c.JSON(names[i-1])
}

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
	newName := new(models.FullNameReq)
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
