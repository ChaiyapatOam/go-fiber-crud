package controller

import (
	"github.com/gofiber/fiber/v2"
)

func CreateCart(c *fiber.Ctx) error {

	return nil
}

func GetCart(c *fiber.Ctx) error {

	c.Status(200).JSON(&fiber.Map{
		"status": "Success",
	})
	return nil
}
