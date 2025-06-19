package controller

import (
	"time"

	"github/Babe-piya/order-management/service"

	"github.com/labstack/echo/v4"
)

type OrderController interface {
	GetOrderByID(c echo.Context) error
	GetOrdersByPagination(c echo.Context) error
	UpdateStatusByID(c echo.Context) error
	CreateOrder(c echo.Context) error
}

type orderController struct {
	Timeout      time.Duration
	OrderService service.OrderService
}

func NewOrderController(timeout time.Duration, orderService service.OrderService) OrderController {
	return &orderController{
		Timeout:      timeout,
		OrderService: orderService,
	}
}
