package models

type Customer struct {
    ID      int
    Name    string
    Email   string
    Address string
}

type Product struct {
    ID          int
    Name        string
    Price       float64
    StockCount  int
}

type Order struct {
    ID         int
    CustomerID int
    OrderDate  string
    Status     string
}

type OrderItem struct {
    ID        int
    OrderID   int
    ProductID int
    Quantity  int
}
