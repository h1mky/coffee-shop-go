package db

import (
	"context"
	"time"
)

func DeleteCard(ctx context.Context, id int) error {

	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()

	if _, err := DB.ExecContext(ctx, "DELETE  FROM coffee_card WHERE id = $1 ", id); err != nil {
		return err
	}
	return nil
}
