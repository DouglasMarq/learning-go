package main

import (
	"fmt"
	"learning-go/src/database"
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

type testStruct struct {
	username string
	email    string
	password string
}

func main() {

	var test = database.Connection()

	testUser := testStruct{"firstUser", "douglas.marq.alves@outlook.com", "555555666555"}

	//testing
	collection := test.Database("Cluster0").Collection("users")

	res, err := collection.Find(nil, testUser)

	if err != nil {
		log.Fatal(err)
	} else {
		for res.Next(nil) {
			var elem testStruct
			err := res.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(elem)
		}
	}

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork:       *prod,
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

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"message": "Route doesn't exist."})
	})

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
