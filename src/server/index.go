package server

import (
	"learning-go/src/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// StartServer Inicia o srevidor Web
func StartServer(){
    app := fiber.New()

	app.Use(cors.New())

	router.AuthRouter(app)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Listen(":3000")
}
