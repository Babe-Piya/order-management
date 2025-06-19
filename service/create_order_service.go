package service

import (
	"context"
	"errors"
	"log/slog"
	"sync"

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
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (srv *orderService) CreateOrder(ctx context.Context, req CreateOrderRequest) (*CreateOrderResponse, error) {
	var wg sync.WaitGroup
	errChan := make(chan error, len(req.OrderDetail))
	for _, order := range req.OrderDetail {
		wg.Add(1)

		go func(order OrderDetail) {
			defer wg.Done()
			tx, err := srv.OrderRepo.BeginTransaction(ctx)
			if err != nil {
				slog.Error("begin transaction error: ", err)
				errChan <- err
			}

			defer func() {
				if err = srv.OrderRepo.RollbackTransaction(ctx, tx); err != nil {
					slog.Error("rollback err: ", err)
				}
			}()

			var totalAmount float64
			var orderItems []repositories.OrderItem
			for _, data := range order.OrderItems {
				totalAmount = totalAmount + (data.Price * float64(data.Quantity))
				orderItems = append(orderItems, repositories.OrderItem{
					ProductName: data.ProductName,
					Quantity:    data.Quantity,
					Price:       data.Price,
				})
			}
			orderID, err := srv.OrderRepo.CreateOrder(ctx, repositories.Order{
				CustomerName: order.CustomerName,
				TotalAmount:  totalAmount,
				Status:       "ORDER CREATED",
			}, tx)
			if err != nil {
				slog.Error("create order error: ", err)
				errChan <- err
			}

			err = srv.OrderRepo.CreateOrderItem(ctx, orderItems, orderID, tx)
			if err != nil {
				slog.Error("create order item error: ", err)
				errChan <- err
			}

			if err = srv.OrderRepo.CommitTransaction(ctx, tx); err != nil {
				errChan <- err
			}
		}(order)
	}
	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, errors.New("insert error")
	}

	return &CreateOrderResponse{
		Code:    "1",
		Message: "Success",
	}, nil
}
