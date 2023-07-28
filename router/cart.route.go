package router

import (
	cartController "github.com/chaiyapatoam/go-fiber-crud/controller"
	"github.com/chaiyapatoam/go-fiber-crud/middleware"
	"github.com/gofiber/fiber/v2"
)

func CartRoute(app *fiber.App) {
	// Middleware
	api := app.Group("/api/cart").Use(middleware.Auth)

	// routes
	api.Post("/", cartController.CreateCart)
	api.Get("/:id", cartController.GetCart)
}