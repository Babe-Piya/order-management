package controller

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github/Babe-piya/order-management/appconstant"
	"github/Babe-piya/order-management/service"

	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) UpdateStatusByID(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, cancel := context.WithTimeout(ctx, ctrl.Timeout*time.Second)
	defer cancel()

	id := c.Param("order_id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("convert order_id type from string to int error:", err)

		return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
			Code:    appconstant.ErrorCode,
			Message: err.Error(),
		})
	}

	req := service.UpdateStatusByIDRequest{}
	err = c.Bind(&req)
	if err != nil {
		slog.Error("bind request to struct error:", err)

		return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
			Code:    appconstant.ErrorCode,
			Message: err.Error(),
		})
	}
	req.ID = int64(orderID)
	resp, err := ctrl.OrderService.UpdateStatusByID(ctx, req)
	if err != nil {
		slog.Error("UpdateStatusByID error:", err)
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
