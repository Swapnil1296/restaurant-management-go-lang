package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// OpenCollection is a function that returns a collection from the database

const dbName = "Restaurant-Management"
const collName = "board"

// var collection *mongo.Collection
// const connectionString = "mongodb+srv://Golang:golang@cluster0.2pndy.mongodb.net/"

// //connect to MongoDB
// func DbInstance() {

// 	// set client options
// 	clientOptions := options.Client().ApplyURI(connectionString)

// 	// connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(),clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connected to MongoDB!")
// 	collection = client.Database(dbName).Collection(collName)
// 	fmt.Println("Collection instance created!")
// }

func DbInstance() *mongo.Client{

	 const connectionString = "mongodb+srv://Golang:golang@cluster0.2pndy.mongodb.net/"

	 fmt.Println("Connecting to MongoDB", connectionString)

	 client, err := mongo.Connect(options.Client().ApplyURI(connectionString))

	 if err != nil {
		 log.Fatal(err)
	 }
	 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	 defer cancel()

	 err = client.Connect(ctx)

	 if err != nil {
		 log.Fatal(err)
	 }

	 fmt.Println("Connected to MongoDB!")
	 return client

}

var Client *mongo.Client = DbInstance()

func OpenCollection(client *mongo.Client, dbName string, collName string) *mongo.Collection {
	return client.Database(dbName).Collection(collName)
}

// package database

// import (
//     "context"
//     "go.mongodb.org/mongo-driver/mongo"
//     "go.mongodb.org/mongo-driver/mongo/options"
// )

// OpenCollection connects to a MongoDB collection
// func OpenCollection(client *mongo.Client, ctx context.Context, collectionName string) *mongo.Collection {
//     db := client.Database("yourDatabaseName") // Replace with your DB name
//     return db.Collection(collectionName)
// }
