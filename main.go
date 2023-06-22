package main

import (
	"ahmadhabibi7159_ToDoList/handlers/auth"
	"ahmadhabibi7159_ToDoList/handlers/todo"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Uncomment this for production (Railway.app will automatically create PORT variable)
	// port := os.Getenv("PORT")
	port := "3000" // Comment this for production
	app := fiber.New()
	app.Static("/", "./client")
	api := app.Group("/api") // All Backend services in /api endpoints
	api.Post("/login", auth.Login)
	api.Get("/todo/list", todo.Lists)

	app.Listen("0.0.0.0:" + port)
}
