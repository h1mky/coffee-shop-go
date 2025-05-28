package main

import (
	"coffee-shop/db"
	"coffee-shop/internal/router"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found, continuing...")
	}

	db.Connect()

	router.StartServer()
}
