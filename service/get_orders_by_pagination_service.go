package service

import (
	"context"
	"log"
	"math"
)

type OrdersByPaginationResponse struct {
	Code             string      `json:"code"`
	Message          string      `json:"message"`
	Data             []OrderData `json:"data"`
	TotalOrder       int         `json:"total_order"`
	Page             int         `json:"page"`
	TotalPage        int         `json:"total_page"`
	TotalOrderInPage int         `json:"total_order_in_page"`
}

func (srv *orderService) GetOrdersByPagination(ctx context.Context, page int, rowOfPage int) (*OrdersByPaginationResponse, error) {
	orders, err := srv.OrderRepo.GetOrdersByPagination(ctx, page, rowOfPage, nil)
	if err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	var orderResults []OrderData
	for _, order := range orders {
		var items []OrderItemData
		for _, orderItem := range order.OrderItems {
			items = append(items, OrderItemData{
				ID:          orderItem.ID,
				ProductName: orderItem.ProductName,
				Price:       orderItem.Price,
				Quantity:    orderItem.Quantity,
			})
		}
		orderResults = append(orderResults, OrderData{
			ID:           order.ID,
			CustomerName: order.CustomerName,
			TotalAmount:  order.TotalAmount,
			Status:       order.Status,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			OrderItems:   items,
		})
	}

	count, err := srv.OrderRepo.GetCountOrder(ctx, nil)
	if err != nil {
		log.Printf(err.Error())

		return nil, err
	}

	totalPage := int(math.Ceil(float64(count) / float64(rowOfPage)))

	return &OrdersByPaginationResponse{
		Code:             "1",
		Message:          "Success",
		Data:             orderResults,
		TotalOrder:       count,
		Page:             page,
		TotalPage:        totalPage,
		TotalOrderInPage: len(orders),
	}, nil
}
