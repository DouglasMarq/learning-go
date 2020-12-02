package server

import (
	"learning-go/src/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet"
)

// StartServer Inicia o srevidor Web
func StartServer(){
    app := fiber.New(fiber.Config{
		Prefork: true,
		DisableStartupMessage: true,
		CaseSensitive: true,
	})

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET,PUT,POST,DELETE,OPTIONS",
		ExposeHeaders: "Content-Type,Authorization,Accept",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	router.AuthRouter(app)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Listen(":3000")
}
