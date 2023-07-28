package router

import (
	// "github.com/chaiyapatoam/go-fiber-crud/controller"
	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App) {
	// api := app.Group("/api")  

	// routes
	AuthRoute(app)
	ProductRoute(app)
	CartRoute(app)
}