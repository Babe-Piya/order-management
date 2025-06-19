package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *orderRepository) GetOrderByID(ctx context.Context, id int64, tx pgx.Tx) (*Order, error) {
	if tx == nil {
		tx, _ = repo.DB.Begin(ctx)
	}

	rows, err := tx.Query(ctx,
		"SELECT orders.*, oi.id AS order_item_id, oi.product_name, oi.quantity, oi.price  FROM orders "+
			"LEFT JOIN order_items oi ON orders.id = oi.order_id WHERE orders.id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orderWithOrderItem, err := pgx.CollectRows(rows, pgx.RowToStructByName[OrderWithOrderItem])
	if err != nil {
		return nil, err
	}

	if len(orderWithOrderItem) == 0 {
		return nil, pgx.ErrNoRows
	}

	result := &Order{
		ID:           orderWithOrderItem[0].ID,
		CustomerName: orderWithOrderItem[0].CustomerName,
		TotalAmount:  orderWithOrderItem[0].TotalAmount,
		Status:       orderWithOrderItem[0].Status,
		CreatedAt:    orderWithOrderItem[0].CreatedAt,
		UpdatedAt:    orderWithOrderItem[0].UpdatedAt,
		OrderItems:   []OrderItem{},
	}

	for _, resultRow := range orderWithOrderItem {
		if resultRow.OrderItemID != nil {
			result.OrderItems = append(result.OrderItems, OrderItem{
				ID:          *resultRow.OrderItemID,
				ProductName: resultRow.ProductName,
				Quantity:    resultRow.Quantity,
				Price:       resultRow.Price,
			})
		}
	}

	return result, nil
}
