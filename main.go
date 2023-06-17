package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
}

var userP = User{
	Username: "ahmadhabibi14",
}

func main() {
	// port := os.Getenv("PORT")
	port := "3000"
	app := fiber.New()

	app.Static("/", "./")

	app.Get("/api/data", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.Set("Access-Control-Allow-Origin", "*")
		return c.JSON(userP)
	})

	url := fmt.Sprintf("0.0.0.0:%s", port)
	app.Listen(url)
}
