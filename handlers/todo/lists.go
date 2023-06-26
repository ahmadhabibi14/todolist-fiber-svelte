package todo

import (
	"ahmadhabibi7159_ToDoList/models"
	"strconv"

	"sort"

	"github.com/gofiber/fiber/v2"
)

func Lists(c *fiber.Ctx) error {
	// Get the page and limit values from the query parameters
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	// Convert the page and limit values to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page value",
		})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid limit value",
		})
	}

	// Calculate the start and end indices for the page
	start := (page - 1) * limit
	end := start + limit

	if start >= len(models.Todos) {
		return c.Status(fiber.StatusOK).JSON(models.Todos)
	}
	if end > len(models.Todos) {
		// Adjust the end index if it exceeds the length of the array
		end = len(models.Todos)
	}

	// Sort By Updated At
	sort.Sort(models.ByUpdatedAt(models.Todos))

	// Get the Todos for the requested page
	pagedTodos := models.Todos[start:end]
	return c.Status(fiber.StatusOK).JSON(pagedTodos)
}
