package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("ecommerce").Collection("customers")
	fmt.Println("Connected to MongoDB.")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var customer struct {
			ID   string `bson:"_id"`
			Name string `bson:"name"`
		}
		cursor.Decode(&customer)
		fmt.Printf("Customer ID: %s, Name: %s\n", customer.ID, customer.Name)
	}
}
