package controller

import (
	"database/sql"

	"api.fiber.practice/repository"
	"github.com/gofiber/fiber/v2"
)

func HandlerGetNames(c *fiber.Ctx) error {
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

	ret, err := r.DBGetAllName()
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"code":   401,
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
