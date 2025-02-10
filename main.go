package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"rahul/db"
	"rahul/utils"
	"strings"
	"time"
)

func main() {
	log.Println("Initializing database connection...")
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Database initialization failed:", err)
	}
	defer database.Close()
	log.Println("Database connection established.")

	for {
		fmt.Println("\nE-commerce Database Query System")
		fmt.Println("1. View Today's Orders")
		fmt.Println("2. View Customer Orders")
		fmt.Println("3. View Product Stock")
		fmt.Println("4. Load Data from CSV")
		fmt.Println("5. View Customer Order Status")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)
		bufio.NewReader(os.Stdin).ReadString('\n')

		log.Printf("User selected option: %d", choice)

		switch choice {
		case 1:
			viewTodaysOrders(database)
		case 2:
			viewCustomerOrders(database)
		case 3:
			viewProductStock(database)
		case 4:
			loadAllCSVData(database)
		case 5:
			viewCustomerOrderStatus(database)
		case 6:
			log.Println("Exiting application.")
			fmt.Println("Exiting...")
			return
		default:
			log.Println("Invalid choice entered.")
			fmt.Println("Invalid choice")
		}
	}
}

func viewTodaysOrders(db *sql.DB) {
	log.Println("Fetching today's orders...")
	today := time.Now().Format("2006-01-02")
	rows, err := db.Query(`
        SELECT o.id, c.name, o.order_date, o.status, p.name, oi.quantity
        FROM orders o 
        JOIN customers c ON o.customer_id = c.id 
        JOIN order_items oi ON o.id = oi.order_id
        JOIN products p ON oi.product_id = p.id
        WHERE date(o.order_date) = ?`, today)
	if err != nil {
		log.Println("Error querying today's orders:", err)
		return
	}
	defer rows.Close()

	log.Println("Displaying today's orders.")
	fmt.Println("\nToday's Orders:")
	fmt.Printf("%-10s %-20s %-20s %-10s %-20s %-10s\n",
		"Order ID", "Customer", "Date", "Status", "Product", "Quantity")
	fmt.Println(strings.Repeat("-", 90))

	for rows.Next() {
		var id int
		var customerName, orderDate, status, productName string
		var quantity int
		err := rows.Scan(&id, &customerName, &orderDate, &status, &productName, &quantity)
		if err != nil {
			log.Println("Error scanning order data:", err)
			continue
		}
		fmt.Printf("%-10d %-20s %-20s %-10s %-20s %-10d\n", id, customerName, orderDate, status, productName, quantity)
	}
}

func viewCustomerOrders(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter customer email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	log.Printf("Fetching orders for customer: %s", email)

	rows, err := db.Query(`
        SELECT o.id, o.order_date, o.status, p.name, p.price, oi.quantity
        FROM orders o
        JOIN customers c ON o.customer_id = c.id
        JOIN order_items oi ON o.id = oi.order_id
        JOIN products p ON oi.product_id = p.id
        WHERE c.email = ?`, email)
	if err != nil {
		log.Println("Error querying customer orders:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nCustomer Orders:")
	fmt.Printf("%-10s %-20s %-10s %-20s %-15s %-10s\n", "Order ID", "Date", "Status", "Product", "Price", "Quantity")
	fmt.Println(strings.Repeat("-", 85))

	for rows.Next() {
		var id int
		var orderDate, status, productName string
		var price float64
		var quantity int
		err := rows.Scan(&id, &orderDate, &status, &productName, &price, &quantity)
		if err != nil {
			log.Println("Error scanning customer order data:", err)
			continue
		}
		fmt.Printf("%-10d %-20s %-10s %-20s ₹%-14.2f %-10d\n", id, orderDate, status, productName, price, quantity)
	}
}

func viewProductStock(db *sql.DB) {
	log.Println("Fetching product stock data...")
	rows, err := db.Query("SELECT id, name, price, stock_count FROM products ORDER BY id")
	if err != nil {
		log.Println("Error querying product stock:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nProduct Stock:")
	fmt.Printf("%-10s %-30s %-15s %-10s\n", "ID", "Product", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 65))

	for rows.Next() {
		var id, stock int
		var name string
		var price float64
		err := rows.Scan(&id, &name, &price, &stock)
		if err != nil {
			log.Println("Error scanning product stock data:", err)
			continue
		}
		fmt.Printf("%-10d %-30s ₹%-14.2f %-10d\n", id, name, price, stock)
	}
}

func viewCustomerOrderStatus(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter customer email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	log.Printf("Fetching order status for customer: %s", email)

	rows, err := db.Query(`
        SELECT o.id, o.order_date, o.status
        FROM orders o
        JOIN customers c ON o.customer_id = c.id
        WHERE c.email = ?`, email)
	if err != nil {
		log.Println("Error querying order status:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nCustomer Order Status:")
	fmt.Printf("%-10s %-20s %-15s\n", "Order ID", "Order Date", "Status")
	fmt.Println(strings.Repeat("-", 50))

	for rows.Next() {
		var id int
		var orderDate, status string
		err := rows.Scan(&id, &orderDate, &status)
		if err != nil {
			log.Println("Error scanning order status data:", err)
			continue
		}
		fmt.Printf("%-10d %-20s %-15s\n", id, orderDate, status)
	}
}

func loadAllCSVData(db *sql.DB) {
	log.Println("Starting to load data from CSV files...")
	fmt.Println("\nLoading data from CSV files...")

	log.Println("Clearing existing data from tables...")
	db.Exec("DELETE FROM order_items")
	db.Exec("DELETE FROM orders")
	db.Exec("DELETE FROM products")
	db.Exec("DELETE FROM customers")
	db.Exec("DELETE FROM sqlite_sequence")

	files := map[string]func(*sql.DB, string) error{
		"data/customers.csv":   utils.LoadCustomersFromCSV,
		"data/products.csv":    utils.LoadProductsFromCSV,
		"data/orders.csv":      utils.LoadOrdersFromCSV,
		"data/order_items.csv": utils.LoadOrderItemsFromCSV,
	}

	for file, loader := range files {
		log.Printf("Loading file: %s", file)
		err := loader(db, file)
		if err != nil {
			log.Printf("Error loading %s: %v", file, err)
		}
	}

	log.Println("CSV data loading completed.")
	fmt.Println("Data loading completed!")
}
