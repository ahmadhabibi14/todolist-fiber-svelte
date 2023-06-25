package todo

import (
	"time"

	"ahmadhabibi7159_ToDoList/models"

	"github.com/gofiber/fiber/v2"
)

type text_store struct {
	Text string `json:"text"`
}

func Add(c *fiber.Ctx) error {
	session_id := c.Get("X-Session-ID")
	textStore := new(text_store)

	if err := c.BodyParser(textStore); err != nil {
		c.Status(fiber.StatusBadRequest)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}

	var todoData = models.Todo{
		Text:       textStore.Text,
		Created_At: time.Now().UTC(),
		Updated_At: time.Now().UTC(),
		Username:   models.Sessions[session_id].Username,
		UserId:     models.Sessions[session_id].UserId,
	}
	models.Todos = append(models.Todos, todoData)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Add item success",
	})
}
