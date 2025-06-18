package service

import (
	"context"
	"log"
	"time"
)

type GetOrderByIDResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    OrderData `json:"data"`
}

type OrderData struct {
	ID           int64
	CustomerName string
	TotalAmount  float64
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (srv *orderService) GetOrderByID(ctx context.Context, id int64) (GetOrderByIDResponse, error) {
	order, err := srv.OrderRepo.GetOrderByID(ctx, id, nil)
	if err != nil {
		log.Println(err)

		return GetOrderByIDResponse{}, err
	}

	return GetOrderByIDResponse{
		Code:    "1",
		Message: "Success",
		Data: OrderData{
			ID:           order.ID,
			CustomerName: order.CustomerName,
			TotalAmount:  order.TotalAmount,
			Status:       order.Status,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
		},
	}, nil
}
