package utils

import (
    "encoding/csv"
    "database/sql"
    "os"
    "strconv"
    "log"
)

func LoadCustomersFromCSV(db *sql.DB, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    stmt, err := db.Prepare("INSERT INTO customers (id, name, email, address) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    for _, record := range records[1:] { // Skip header
        id, _ := strconv.Atoi(record[0])
        _, err := stmt.Exec(id, record[1], record[2], record[3])
        if err != nil {
            log.Printf("Error inserting customer: %v", err)
        }
    }
    return nil
}

func LoadProductsFromCSV(db *sql.DB, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    stmt, err := db.Prepare("INSERT INTO products (id, name, price, stock_count) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    for _, record := range records[1:] {
        id, _ := strconv.Atoi(record[0])
        price, _ := strconv.ParseFloat(record[2], 64)
        stock, _ := strconv.Atoi(record[3])
        _, err := stmt.Exec(id, record[1], price, stock)
        if err != nil {
            log.Printf("Error inserting product: %v", err)
        }
    }
    return nil
}

func LoadOrdersFromCSV(db *sql.DB, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    stmt, err := db.Prepare("INSERT INTO orders (id, customer_id, order_date, status) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    for _, record := range records[1:] {
        id, _ := strconv.Atoi(record[0])
        customerID, _ := strconv.Atoi(record[1])
        _, err := stmt.Exec(id, customerID, record[2], record[3])
        if err != nil {
            log.Printf("Error inserting order: %v", err)
        }
    }
    return nil
}

func LoadOrderItemsFromCSV(db *sql.DB, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    stmt, err := db.Prepare("INSERT INTO order_items (id, order_id, product_id, quantity) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    for _, record := range records[1:] {
        id, _ := strconv.Atoi(record[0])
        orderID, _ := strconv.Atoi(record[1])
        productID, _ := strconv.Atoi(record[2])
        quantity, _ := strconv.Atoi(record[3])
        _, err := stmt.Exec(id, orderID, productID, quantity)
        if err != nil {
            log.Printf("Error inserting order item: %v", err)
        }
    }
    return nil
}
