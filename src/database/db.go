package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type MongoDatastore struct {
// 	db      *mongo.Database
// 	Session *mongo.Client
// }

//Connection conexão do mongo
func Connection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:da123456789@cluster0.twaan.gcp.mongodb.net/Cluster0?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
