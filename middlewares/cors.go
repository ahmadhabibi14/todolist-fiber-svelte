package middlewares

import "github.com/gofiber/fiber/v2"

func CORS(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-Control, X-Requested-With, X-Session-ID")
	return c.Next()
}
