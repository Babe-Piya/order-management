package service

import (
	"context"
	"log"

	"github/Babe-piya/order-management/repositories"
)

type CreateOrderRequest struct {
	OrderDetail []OrderDetail `json:"order_detail"`
}

type OrderDetail struct {
	CustomerName string                `json:"customer_name"`
	OrderItems   []CreateOrderItemData `json:"order_items"`
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

	var totalAmount float64
	var orderItems []repositories.OrderItem
	for _, data := range req.OrderDetail[0].OrderItems {
		totalAmount = totalAmount + (data.Price * float64(data.Quantity))
		orderItems = append(orderItems, repositories.OrderItem{
			ProductName: data.ProductName,
			Quantity:    data.Quantity,
			Price:       data.Price,
		})
	}
	orderID, err := srv.OrderRepo.CreateOrder(ctx, repositories.Order{
		CustomerName: req.OrderDetail[0].CustomerName,
		TotalAmount:  totalAmount,
		Status:       "ORDER CREATED",
	}, tx)
	if err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	err = srv.OrderRepo.CreateOrderItem(ctx, orderItems, orderID, tx)
	if err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	if err = srv.OrderRepo.CommitTransaction(ctx, tx); err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	return &CreateOrderResponse{
		Code:        "1",
		Message:     "Success",
		TotalAmount: totalAmount,
	}, nil
}
