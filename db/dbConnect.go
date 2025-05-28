package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
)

var DB *sqlx.DB

func Connect() {
	connect := os.Getenv("DB_CONNECT")
	if connect == "" {
		log.Fatal("DB_CONNECT environment variable is not set")
	}

	// Добавляем SSL параметры если их нет
	if !strings.Contains(connect, "sslmode") {
		if strings.Contains(connect, "?") {
			connect += "&sslmode=require"
		} else {
			connect += "?sslmode=require"
		}
	}

	log.Printf("Attempting to connect to database...")

	var err error
	DB, err = sqlx.Connect("postgres", connect)
	if err != nil {
		log.Printf("Connection string: %s", connect)
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Проверяем соединение
	if err := DB.Ping(); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Successfully connected to database")
}
