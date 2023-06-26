package todo

import (
	"ahmadhabibi7159_ToDoList/models"

	"github.com/gofiber/fiber/v2"
)

func Lists(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(models.Todos)
}
