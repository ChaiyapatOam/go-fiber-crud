package router

import (
	"github.com/chaiyapatoam/go-fiber-crud/controller"
	"github.com/chaiyapatoam/go-fiber-crud/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProductRoute(app *fiber.App) {
	// Middleware
	api := app.Group("/api/product").Use(middleware.Auth)

	api.Get("/", controller.GetAllProducts)
	api.Get("/:id", controller.GetProductById)
	api.Post("/", controller.AddProduct)
	api.Patch("/:id", controller.EditProduct)
	api.Delete("/:id", controller.DeleteProduct)
}
