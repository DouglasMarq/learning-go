package main

import (
	"context"
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
	"go.mongodb.org/mongo-driver/bson"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", true, "Enable prefork in Production")
)

type userStruct struct {
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

func main() {
	var connection = database.Connection()

	user := userStruct{"IDPBBrisa", "douglas.marq.alves@outlook.com", "1234"}
	var foundUser userStruct

	//testing
	collection := connection.Database("Cluster0").Collection("users")

	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&foundUser)

	if err != nil {
		result, err := collection.InsertOne(context.Background(), bson.M{"username": user.Username, "password": user.Password, "email": user.Email})
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		}
		if error := collection.FindOne(context.Background(), bson.M{"_id": result.InsertedID}).Decode(&foundUser); error != nil {
			log.Fatal(error)
		}
	}
	fmt.Println(foundUser)

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
