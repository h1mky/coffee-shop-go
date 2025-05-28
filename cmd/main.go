package main

import (
	"coffee-shop/db"
	"coffee-shop/internal/router"
)

func main() {

	db.Connect()

	router.StartServer()
}
