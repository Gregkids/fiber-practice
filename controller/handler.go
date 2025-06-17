package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func handlerGet(c *fiber.Ctx) error {
	fmt.Fprintf(c, "%s\n", c.Params("Name"))
	fmt.Fprintf(c, "%s\n", c.Params("Age"))

	return nil
}
