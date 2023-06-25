package global

import (
	"ahmadhabibi7159_ToDoList/models"

	"github.com/gofiber/fiber/v2"
)

type stats struct {
	LoggedIn       int `json:"logged_in"`
	UserRegistered int `json:"user_registered"`
	TodoCreated    int `json:"todo_created"`
}

func Stats(c *fiber.Ctx) error {
	statsData := stats{
		LoggedIn:       int(models.TotalLoggedIn),
		UserRegistered: len(models.UserStores),
		TodoCreated:    len(models.Todos),
	}

	return c.Status(fiber.StatusOK).JSON(statsData)
}
