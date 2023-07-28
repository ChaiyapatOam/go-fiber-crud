package router

import (
	"github.com/chaiyapatoam/go-fiber-crud/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	// Middleware
	api := app.Group("/api/auth")  
	
	// routes
	api.Post("/login", controller.Login)
	// api.Post("/signup", controller.SignUp) 
}