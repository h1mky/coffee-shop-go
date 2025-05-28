package router

import (
	http2 "coffee-shop/internal/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
)

func StartServer() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

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
	fmt.Printf("Server started on :%s\n", port)

	r.Get("/cards/{id}", http2.GetSingleCardHandler)

	r.Post("/cards", http2.PostNewCardHandler)

	r.Delete("/cards/{id}", http2.DeleteCardHandler)

	r.Patch("/cards/{id}", http2.UpdateCardHandle)

	r.Post("/admin", http2.AdminPanelRedirect)

	log.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
