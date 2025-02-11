package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:secret@localhost:5432/mydb")
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected to PostgreSQL using pgx.")

	rows, err := conn.Query(context.Background(), "SELECT id, name FROM customers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("Customer ID: %d, Name: %s\n", id, name)
	}
}
