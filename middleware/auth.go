package middleware

import (
	"time"

	"github.com/chaiyapatoam/go-fiber-crud/service"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(401).JSON(&fiber.Map{
			"message": "No Token Provided",
			"success": false,
		})
	}

	// if _, err := service.ValidateToken(token); err != nil {
	// 	c.Status(401).JSON(&fiber.Map{
	// 		"message": "Unauthorized",
	// 		"success": false,
	// 	})
	// }

	validToken, err := service.ValidateToken(token)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{"success": false})
	}

	if validToken == nil {
		return c.Status(401).JSON(&fiber.Map{
			"message": "Unauthorized",
			"success": false,
		})
	}
	// Check Token expired
	now := time.Now().Unix()

	if now > validToken.Expires {
		// Return status 401 and Token Expired.
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized Token Expired",
		})
	}
	return c.Next()
}
