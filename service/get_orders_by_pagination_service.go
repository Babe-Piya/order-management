package service

import (
	"context"
	"log/slog"
	"math"

	"github/Babe-piya/order-management/appconstant"
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

	orders, err := srv.OrderRepo.GetOrdersByPagination(ctx, page, rowOfPage, tx)
	if err != nil {
		slog.Error("GetOrdersByPagination error:", err)

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

	count, err := srv.OrderRepo.GetCountOrder(ctx, tx)
	if err != nil {
		slog.Error("GetCountOrder error:", err)

		return nil, err
	}

	err = srv.OrderRepo.CommitTransaction(ctx, tx)
	if err != nil {
		slog.Error("commit transaction error:", err)

		return nil, err
	}

	totalPage := int(math.Ceil(float64(count) / float64(rowOfPage)))

	return &OrdersByPaginationResponse{
		Code:             appconstant.SuccessCode,
		Message:          appconstant.SuccessMsg,
		Data:             orderResults,
		TotalOrder:       count,
		Page:             page,
		TotalPage:        totalPage,
		TotalOrderInPage: len(orders),
	}, nil
}
