package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Pagination struct {
	Page      int
	RowOfPage int
}

func (repo *orderRepository) GetOrdersByPagination(ctx context.Context, page int, rowOfPage int, tx pgx.Tx) ([]Order, error) {
	if tx == nil {
		tx, _ = repo.DB.BeginTx(ctx, pgx.TxOptions{})
	}

	offset := (page - 1) * rowOfPage
	rows, err := tx.Query(ctx,
		"SELECT o.*, oi.id AS order_item_id, oi.product_name, oi.quantity, oi.price "+
			"FROM (SELECT * FROM orders ORDER BY orders.id OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY) AS o "+
			"LEFT JOIN order_items oi ON o.id = oi.order_id ", offset, rowOfPage)
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

	orderMap := make(map[int64]*Order)
	for _, resultRow := range orderWithOrderItem {
		order, ok := orderMap[resultRow.ID]
		if !ok {
			order = &Order{
				ID:           resultRow.ID,
				CustomerName: resultRow.CustomerName,
				TotalAmount:  resultRow.TotalAmount,
				Status:       resultRow.Status,
				CreatedAt:    resultRow.CreatedAt,
				UpdatedAt:    resultRow.UpdatedAt,
				OrderItems:   []OrderItem{},
			}
			orderMap[resultRow.ID] = order
		}

		if resultRow.OrderItemID != nil {
			order.OrderItems = append(order.OrderItems, OrderItem{
				ID:          *resultRow.OrderItemID,
				ProductName: resultRow.ProductName,
				Quantity:    resultRow.Quantity,
				Price:       resultRow.Price,
			})
		}
	}

	var result []Order
	for _, u := range orderMap {
		result = append(result, *u)
	}

	return result, nil
}
