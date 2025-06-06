package db

import (
	"coffee-shop/models"
	"context"
	"log"
	"time"
)

func GetAllCard(ctx context.Context) ([]models.CoffeeCard, error) {

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)

	defer cancel()

	var products []models.CoffeeCard

	err := DB.SelectContext(ctx, &products, "SELECT * FROM coffee_card ORDER BY created_at DESC")

	if err != nil {
		log.Println("Error fetching a card", err)
		return nil, err
	}
	return products, nil
}
