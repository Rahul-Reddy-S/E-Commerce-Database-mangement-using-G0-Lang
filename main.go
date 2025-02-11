package main

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=secret dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	err = db.Select(&users, "SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println("User ID:", user.ID, "Name:", user.Name)
	}
}
