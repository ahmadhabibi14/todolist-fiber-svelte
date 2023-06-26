package todo

import (
	"ahmadhabibi7159_ToDoList/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Overwrite(c *fiber.Ctx) error {
	session_id := c.Get("X-Session-ID")
	user_id := models.Sessions[session_id].UserId
	textStore := new(models.EditTextStore)

	if err := c.BodyParser(textStore); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}

	for i := range models.Todos {
		if models.Todos[i].UserId == user_id && models.Todos[i].Id == textStore.Id {
			models.Todos[i].Text = textStore.Text
			models.Todos[i].Updated_At = time.Now().UTC()
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Edit item success",
	})
}
