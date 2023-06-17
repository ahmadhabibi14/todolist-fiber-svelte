package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
}

var userP = User{
	Username: "ahmadhabibi14",
}

func main() {
	Port := os.Getenv("PORT")
	app := fiber.New()

	app.Get("/api/data", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.Set("Access-Control-Allow-Origin", "*")
		return c.JSON(userP)
	})

	url := fmt.Sprintf("0.0.0.0:%s", Port)
	app.Listen(url)
}
