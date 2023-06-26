package todo

import (
	"ahmadhabibi7159_ToDoList/models"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	session_id := c.Get("X-Session-ID")
	user_id := models.Sessions[session_id].UserId
	deleteTodoStore := new(models.DeleteTodoStore)

	if err := c.BodyParser(deleteTodoStore); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}

	for i := range models.Todos {
		if models.Todos[i].UserId == user_id && models.Todos[i].Id == deleteTodoStore.Id {
			models.Todos = models.DeleteTodo(models.Todos, i)
			models.TotalTodoDeleted++
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete item success",
	})
}
