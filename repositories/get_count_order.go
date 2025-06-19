package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *orderRepository) GetCountOrder(ctx context.Context, tx pgx.Tx) (int, error) {
	if tx == nil {
		tx, _ = repo.DB.Begin(ctx)
	}

	var count int
	err := tx.QueryRow(ctx, "SELECT COUNT(*)  FROM orders").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
