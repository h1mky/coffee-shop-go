package db

import (
	"context"
	"time"
)

func PostNewCard(ctx context.Context, card interface{}) error {

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	_, err := DB.NamedExec(
		`INSERT INTO coffee_card ( img, coffee_name, country, price, recommended, description)
         VALUES ( :img, :coffee_name, :country, :price, :recommended, :description)`,
		card,
	)
	if err != nil {
		return err

	}
	return nil
}
