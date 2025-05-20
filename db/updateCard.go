package db

import (
	"coffee-shop/models"
	"context"
	"time"
)

func UpdateCard(ctx context.Context, id int, card models.CoffeeCard) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)

	defer cancel()

	query := `
		UPDATE coffee_card 
		SET 
			img = :img,
			coffee_name = :coffee_name,
			country = :country,
			price = :price,
			recommended = :recommended,
			description = :description
		WHERE id = :id
	`

	params := ToParams(id, card)

	if _, err := DB.NamedExecContext(ctx, query, params); err != nil {
		return err
	}
	return nil
}
