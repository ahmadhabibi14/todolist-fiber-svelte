package main

import (
	"ahmadhabibi7159_ToDoList/handlers/auth"
	"ahmadhabibi7159_ToDoList/handlers/global"
	"ahmadhabibi7159_ToDoList/handlers/todo"
	"ahmadhabibi7159_ToDoList/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	// Uncomment this for production (Railway.app will automatically create PORT variable)
	// port := os.Getenv("PORT")
	port := "4000" // Comment this for production
	app := fiber.New()
	app.Use(limiter.New(middlewares.Limiter))
	app.Static("/", "./client")
	api := app.Group("/api") // All Backend services in /api endpoints

	// Auth endpoints
	api.Post("/login", auth.Login)                                                  // Login
	api.Delete("/logout", middlewares.CookieSession, middlewares.CORS, auth.Logout) // Logout

	// To-Do-List endpoints
	api.Get("/todo/list", todo.Lists)
	api.Post("/todo/add", middlewares.CookieSession, middlewares.CORS, todo.Add)
	api.Put("/todo/overwrite", middlewares.CookieSession, middlewares.CORS, todo.Overwrite)
	api.Delete("/todo/delete", middlewares.CookieSession, middlewares.CORS, todo.Delete)

	// Global endpoints
	api.Get("/stats", global.Stats)

	app.Listen("0.0.0.0:" + port)
}
