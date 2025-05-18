package db

import (
	"coffee-shop/models"
	"context"
	"log"
)

func GetAllCard(ctx context.Context) ([]models.CoffeeCard, error) {
	var products []models.CoffeeCard
	err := DB.SelectContext(ctx, &products, "SELECT * FROM coffee_card")
	if err != nil {
		log.Println("Error fetching a card", err)
		return nil, err
	}
	return products, nil
}
