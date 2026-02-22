package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "Learn"
const colName = "course"

var Collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")

}

