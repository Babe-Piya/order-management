package service

import (
	"context"

	"github/Babe-piya/order-management/repositories"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, id int64) (GetOrderByIDResponse, error)
}

type orderService struct {
	OrderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &orderService{
		OrderRepo: orderRepo,
	}
}
