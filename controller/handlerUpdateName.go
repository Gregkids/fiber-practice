package controller

import (
	"database/sql"

	"api.fiber.practice/models"
	"api.fiber.practice/repository"
	"github.com/gofiber/fiber/v2"
)

func HandlerUpdateName(c *fiber.Ctx) error {
	// Connecting to Database
	db, err := sql.Open("postgres", "host=localhost user=postgres password=darageta dbname=postgres sslmode=disable")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":   500,
			"status": "SQL Error",
			"msg":    err.Error(),
		})
	}

	// Parse Database
	r := repository.NameSQL{DB: db}

	// Declare Requests
	reqName := new(models.FullNameReq)
	reqID := c.QueryInt("id")
	err = c.BodyParser(reqName)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "body Error",
			"msg":    err.Error(),
		})
	}

	// Execute Method
	err = r.DBUpdateName(reqName, reqID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":   400,
			"status": "Error",
			"msg":    err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"code":   201,
		"status": "success",
		"msg":    "Name Updated",
	})
}
