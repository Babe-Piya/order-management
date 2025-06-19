package controller

import (
	"github/Babe-piya/order-management/service"

	"github.com/labstack/echo/v4"
)

type OrderController interface {
	GetOrderByID(c echo.Context) error
	GetOrdersByPagination(c echo.Context) error
	UpdateStatusByID(c echo.Context) error
}

type orderController struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{
		OrderService: orderService,
	}
}
