package controller

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github/Babe-piya/order-management/appconstant"

	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) GetOrderByID(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, cancel := context.WithTimeout(ctx, ctrl.Timeout*time.Second)
	defer cancel()

	id := c.Param("order_id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("convert id type from string to int error:", err)

		return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
			Code:    appconstant.ErrorCode,
			Message: err.Error(),
		})
	}

	resp, err := ctrl.OrderService.GetOrderByID(ctx, int64(orderID))
	if err != nil {
		slog.Error("GetOrderByID error:", err)
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

	return c.JSON(http.StatusOK, resp)
}
