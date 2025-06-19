package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *orderRepository) CreateOrder(ctx context.Context, data Order, tx pgx.Tx) (int64, error) {
	if tx == nil {
		tx, _ = repo.DB.BeginTx(ctx, pgx.TxOptions{})
	}

	var orderID int64
	err := tx.QueryRow(ctx,
		"INSERT INTO orders (customer_name, total_amount, status) VALUES ($1, $2, $3) RETURNING id",
		data.CustomerName, data.TotalAmount, data.Status).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}
