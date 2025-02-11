package main

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct {
	ID   int
	Name string
}

func main() {
	dsn := "user=postgres password=secret dbname=mydb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("Connected to PostgreSQL using GORM.")

	var customers []Customer
	db.Find(&customers)

	for _, customer := range customers {
		fmt.Printf("Customer ID: %d, Name: %s\n", customer.ID, customer.Name)
	}
}
