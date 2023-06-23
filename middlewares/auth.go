package middlewares

import (
	"ahmadhabibi7159_ToDoList/models"

	"github.com/gofiber/fiber/v2"
)

func CookieSession(c *fiber.Ctx) error {
	session_id := c.Cookies("session_id")
	if _, ok := models.Sessions[session_id]; !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized Access",
		})
	}
	return c.Next()
}
