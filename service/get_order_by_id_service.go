package service

import (
	"context"
	"log/slog"
	"time"

	"github/Babe-piya/order-management/appconstant"
)

type GetOrderByIDResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    OrderData `json:"data"`
}

type OrderData struct {
	ID           int64           `json:"id"`
	CustomerName string          `json:"customer_name"`
	TotalAmount  float64         `json:"total_amount"`
	Status       string          `json:"status"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	OrderItems   []OrderItemData `json:"order_items"`
}
type OrderItemData struct {
	ID          int64   `json:"id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (srv *orderService) GetOrderByID(ctx context.Context, id int64) (*GetOrderByIDResponse, error) {
	tx, err := srv.OrderRepo.BeginTransaction(ctx)
	if err != nil {
		slog.Error("begin transaction error:", err)

		return nil, err
	}

	defer func() {
		err = srv.OrderRepo.RollbackTransaction(ctx, tx)
		if err != nil {
			slog.Warn("rollback transaction error:", err)
		}
	}()

	order, err := srv.OrderRepo.GetOrderByID(ctx, id, tx)
	if err != nil {
		slog.Error("GetOrderByID error:", err)

		return nil, err
	}

	var orderItems []OrderItemData
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, OrderItemData{
			ID:          item.ID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
		})
	}

	err = srv.OrderRepo.CommitTransaction(ctx, tx)
	if err != nil {
		slog.Error("commit transaction error:", err)

		return nil, err
	}

	return &GetOrderByIDResponse{
		Code:    appconstant.SuccessCode,
		Message: appconstant.SuccessMsg,
		Data: OrderData{
			ID:           order.ID,
			CustomerName: order.CustomerName,
			TotalAmount:  order.TotalAmount,
			Status:       order.Status,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			OrderItems:   orderItems,
		},
	}, nil
}
