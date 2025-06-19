package service

import (
	"context"
	"log"

	"github/Babe-piya/order-management/repositories"
)

type CreateOrderRequest struct {
	OrderDeTail []OrderDetail `json:"order_detail"`
}

type OrderDetail struct {
	CustomerName string          `json:"customer_name"`
	TotalAmount  float64         `json:"total_amount"`
	Status       string          `json:"status"`
	OrderItems   []OrderItemData `json:"order_items"`
}

type CreateOrderItemData struct {
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type CreateOrderResponse struct {
	Code        string  `json:"code"`
	Message     string  `json:"message"`
	TotalAmount float64 `json:"total_amount"`
}

func (srv *orderService) CreateOrder(ctx context.Context, req CreateOrderRequest) (*CreateOrderResponse, error) {
	// TODO: use goroutine for many order wrap tx in goroutine and use connection pool
	tx, err := srv.OrderRepo.BeginTransaction(ctx)
	if err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	defer func() {
		if err = srv.OrderRepo.RollbackTransaction(ctx, tx); err != nil {
			log.Printf(err.Error())
		}
	}()

	_, err = srv.OrderRepo.CreateOrder(ctx, repositories.Order{}, tx)
	if err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	if err = srv.OrderRepo.CommitTransaction(ctx, tx); err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	return &CreateOrderResponse{
		Code:    "1",
		Message: "Success",
	}, nil
}
