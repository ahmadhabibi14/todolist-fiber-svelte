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
			// Delete previous session
			for key, value := range models.Sessions {
				if value.Username == creds.Username {
					delete(models.Sessions, key)
				}
			}
			// create a new session
			models.Sessions[sessionToken] = models.Session{
				Username: user.Username,
				Expiry:   expiresAt,
			}
			c.Cookie(&cookie)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"user_id":  user.Id,
				"username": creds.Username,
			})
		}
	}
	userModel.NewUser(*creds)
	// create a new session
	models.Sessions[sessionToken] = models.Session{
		Username: creds.Username,
		Expiry:   expiresAt,
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id":  models.IdAutoInc - 1,
		"username": creds.Username,
	})
}
