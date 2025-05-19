package router

import (
	http2 "coffee-shop/iternal/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func StartServer() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/cards", http2.GetAllCardHandler)
	fmt.Println("Server started on :3000")

	r.Get("/cards/:id", http2.GetSingleCardHandler)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
