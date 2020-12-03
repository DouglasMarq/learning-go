package main

import (
	"learning-go/src/handlers"
	"learning-go/src/router"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", true, "Enable prefork in Production")
)

func main() {
	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod,
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

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	//Auth handling
	router.AuthRouter(app)

	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
