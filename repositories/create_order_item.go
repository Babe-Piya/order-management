package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *orderRepository) CreateOrderItem(ctx context.Context, data []OrderItem, orderID int64, tx pgx.Tx) error {
	if tx == nil {
		tx, _ = repo.DB.BeginTx(ctx, pgx.TxOptions{})
	}

	var rows [][]interface{}
	for _, orderItem := range data {
		rows = append(rows, []interface{}{
			orderID, orderItem.ProductName, orderItem.Quantity, orderItem.Price,
		})
	}

	_, err := tx.CopyFrom(ctx,
		pgx.Identifier{"order_items"},
		[]string{"order_id", "product_name", "quantity", "price"},
		pgx.CopyFromRows(rows))
	if err != nil {
		return err
	}

	return nil
}
