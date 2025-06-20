package controller

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github/Babe-piya/order-management/appconstant"
	"github/Babe-piya/order-management/service"

	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) CreateOrder(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, cancel := context.WithTimeout(ctx, ctrl.Timeout*time.Second)
	defer cancel()

	req := service.CreateOrderRequest{}
	err := c.Bind(&req)
	if err != nil {
		slog.Error("bind request to struct error:", err)

		return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
			Code:    appconstant.ErrorCode,
			Message: err.Error(),
		})
	}

	resp, err := ctrl.OrderService.CreateOrder(ctx, req)
	if err != nil {
		slog.Error("CreateOrder error:", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return c.JSON(http.StatusRequestTimeout, appconstant.ErrorResponse{
				Code:    appconstant.ErrorCode,
				Message: appconstant.ErrorTimeout,
			})
		}

		return c.JSON(http.StatusInternalServerError, appconstant.ErrorResponse{
			Code:    appconstant.ErrorCode,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, resp)
}
