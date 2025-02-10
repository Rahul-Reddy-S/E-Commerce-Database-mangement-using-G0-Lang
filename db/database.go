package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func InitDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./ecommerce.db")
    if err != nil {
        return nil, err
    }

    err = createTables(db)
    if err != nil {
        return nil, err
    }
    
    return db, nil
}

func createTables(db *sql.DB) error {
    tables := []string{
        `CREATE TABLE IF NOT EXISTS customers (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            address TEXT
        );`,
        `CREATE TABLE IF NOT EXISTS products (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            price REAL NOT NULL,
            stock_count INTEGER NOT NULL
        );`,
        `CREATE TABLE IF NOT EXISTS orders (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            customer_id INTEGER,
            order_date TEXT NOT NULL,
            status TEXT NOT NULL,
            FOREIGN KEY(customer_id) REFERENCES customers(id)
        );`,
        `CREATE TABLE IF NOT EXISTS order_items (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            order_id INTEGER,
            product_id INTEGER,
            quantity INTEGER NOT NULL,
            FOREIGN KEY(order_id) REFERENCES orders(id),
            FOREIGN KEY(product_id) REFERENCES products(id)
        );`,
    }

    for _, table := range tables {
        _, err := db.Exec(table)
        if err != nil {
            log.Printf("Error creating table: %v", err)
            return err
        }
    }
    return nil
}
