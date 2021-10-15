package router

import (
	"learning-go/src/handlers"

	"github.com/gofiber/fiber/v2"
)

// AuthRouter is for routing auth handler
func AuthRouter(app *fiber.App) {

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.UserList)
	v1.Post("/users", handlers.UserCreate)
}
