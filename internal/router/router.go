package router

import (
	http2 "coffee-shop/internal/http"
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
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/cards", http2.GetAllCardHandler)
	fmt.Println("Server started on :3000")

	r.Get("/cards/{id}", http2.GetSingleCardHandler)

	r.Post("/cards", http2.PostNewCardHandler)

	r.Delete("/cards/{id}", http2.DeleteCardHandler)

	r.Patch("/cards/{id}", http2.UpdateCardHandle)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
