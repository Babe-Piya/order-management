package controller

import (
	"log"
	"net/http"

	"github/Babe-piya/order-management/appconstant"
	"github/Babe-piya/order-management/service"

	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) CreateOrder(c echo.Context) error {
	ctx := c.Request().Context()

	req := service.CreateOrderRequest{}
	err := c.Bind(&req)
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
			Code:    "0",
			Message: err.Error(),
		})
	}

	resp, err := ctrl.OrderService.CreateOrder(ctx, req)
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusInternalServerError, appconstant.ErrorResponse{
			Code:    "0",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
