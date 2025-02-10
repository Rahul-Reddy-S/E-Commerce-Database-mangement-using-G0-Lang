# E-Commerce Database Manager

## Overview

The **E-Commerce Database Manager** is a command-line tool built using Go (Golang) and SQLite3, designed to manage data for a small e-commerce system. It allows users to track customer orders, view product stock, and load data from CSV files.

## Features

- **Data Loading**: Import customer, product, order, and order item data from CSV files.
- **Order Tracking**: View today's orders and order history for specific customers.
- **Inventory Management**: Display product stock levels with prices.
- **Simple Interface**: Easy-to-use command-line interface with menu options.
- **SQLite Database**: Uses SQLite3 for local data storage.
- **Debug Logging**: Logs operations and errors to a file for debugging.



## Installation

1. **Prerequisites:**
   - Go (Golang) installed (version 1.21 or higher)
   - SQLite3 installed (optional, but recommended for viewing the database directly)

2. **Clone the Repository:**
https://github.com/Rahul-Reddy-S/E-Commerce-Database-mangement-using-G0-Lang.git


3. **Initialize Go Modules and Download Dependencies:**

go mod init ecommerce-db
go mod tidy


4. **Run the Application:**
go run main .go



