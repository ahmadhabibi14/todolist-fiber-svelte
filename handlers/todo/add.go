package todo

import (
	"time"

	"ahmadhabibi7159_ToDoList/models"

	"github.com/rs/xid"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {
	guid := xid.New()
	session_id := c.Get("X-Session-ID")
	textStore := new(models.AddTextStore)

	if err := c.BodyParser(textStore); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}

	var todoData = models.Todo{
		Id:         guid.String(),
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
