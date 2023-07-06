package router

import (
	"github.com/chaiyapatoam/go-fiber-crud/controller"
	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App) {
	// Middleware
	api := app.Group("/api")  
	
	// routes
	api.Get("/", controller.GetAllProducts)
	api.Get("/:id", controller.GetProductById)
	api.Post("/", controller.AddProduct)
	api.Patch("/:id", controller.EditProduct)
	api.Delete("/:id", controller.DeleteProduct)
}