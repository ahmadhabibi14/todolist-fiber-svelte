package auth

import (
	"ahmadhabibi7159_ToDoList/models"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	session_id := c.Get("X-Session-ID")
	// Delete session from sessions map
	for _, value := range models.Sessions {
		username := value.Username
		// Set logged_in to false
		for _, user := range models.UserStores {
			if user.Username == username {
				user.LoggedIn = false
			}
		}
		models.TotalLoggedIn -= 1
		delete(models.Sessions, session_id)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout Successfully !!",
	})
}
