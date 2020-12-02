package router

import (
	controller "learning-go/src/controller"

	"github.com/gofiber/fiber/v2"
)

// AuthRouter is for routing auth handler
func AuthRouter(app *fiber.App) error {

	api := app.Group("/auth")

	// api.Post("/login", controller.LoginHandler)
	api.Post("/register", controller.RegisterController)
	// api.Get("/token", middleware.ProtectedUser, controller.CheckTokenHandler)
	// api.Get("/refresh-token", controller.RequestTokenHandler)

	return nil
}
