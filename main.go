package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"your_project/ent"
	"your_project/ent/customer"
)

func main() {
	client, err := ent.Open("postgres", "user=postgres password=secret dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	defer client.Close()

	ctx := context.Background()
	fmt.Println("Connected to PostgreSQL using Ent.")

	customers, err := client.Customer.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range customers {
		fmt.Printf("Customer ID: %d, Name: %s\n", c.ID, c.Name)
	}
}
