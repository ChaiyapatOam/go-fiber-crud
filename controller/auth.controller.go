package controller

import (
	"github.com/chaiyapatoam/go-fiber-crud/service"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// Simple Login with Gmail

	type B struct {
		Email string `json:"email"`
	}
	// Get email from Body
	b := new(B)
	err := c.BodyParser(b)
	if err != nil {
		return c.SendStatus(500)
	}

	token, err := service.GenerateToken(b.Email)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
		})
	}

	c.Status(200).JSON(&fiber.Map{
		"email" : b.Email,
		"Token": token,
	})
	return nil
}

func SignUp(c *fiber.Ctx) error {
	return nil
}
