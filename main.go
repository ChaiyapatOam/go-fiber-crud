package main

import (
	"fmt"

	"github.com/chaiyapatoam/go-fiber-crud/database"
	"github.com/chaiyapatoam/go-fiber-crud/router"
	"github.com/gofiber/fiber/v2"

)

func main() {
	err := database.Connect()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	router.MainRoute(app)
	app.Listen(":3000")
}
