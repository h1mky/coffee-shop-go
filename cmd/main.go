package main

import (
	"coffee-shop/iternal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	router.StartServer()
}
