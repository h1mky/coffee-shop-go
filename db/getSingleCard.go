package db

import (
	"coffee-shop/models"
	"context"
	"log"
	"time"
)

func GetSingleCard(ctx context.Context, id int) (models.CoffeeCard, error) {

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)

	defer cancel()

	var product models.CoffeeCard

	err := DB.GetContext(ctx, &product, "SELECT * FROM coffee_card WHERE id = $1", id)

	if err != nil {
		log.Println("Error fetching a card", err)
		return product, err
	}
	return product, nil

}
