package auth

import (
	"ahmadhabibi7159_ToDoList/models"
	"time"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	creds := new(models.Credentials)
	userModel := models.User{}

	if err := c.BodyParser(creds); err != nil {
		c.Status(fiber.StatusBadRequest)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	cookie := fiber.Cookie{
		Name:     "session_id",
		Value:    sessionToken,
		Expires:  expiresAt,
		HTTPOnly: false,
	}

	for _, user := range models.UserStores {
		if user.Username == creds.Username && user.Password == creds.Password {
			models.Sessions[sessionToken] = models.Session{
				Username: user.Username,
				Expiry:   expiresAt,
			}
			c.Cookie(&cookie)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Login Successful !!",
			})
		}
	}
	userModel.NewUser(*creds)
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Successful !!",
	})
}
