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

func (ctrl *orderController) GetOrdersByPagination(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, cancel := context.WithTimeout(ctx, ctrl.Timeout*time.Second)
	defer cancel()

	page := 1
	var err error
	pageStr := c.QueryParam("page")
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			slog.Error("convert page type from string to int error:", err)

			return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
				Code:    appconstant.ErrorCode,
				Message: err.Error(),
			})
		}
	}

	rowOfPage := 50
	rowOfPageStr := c.QueryParam("row_of_page")
	if rowOfPageStr != "" {
		rowOfPage, err = strconv.Atoi(rowOfPageStr)
		if err != nil {
			slog.Error("convert row_of_page type from string to int error:", err)

			return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
				Code:    appconstant.ErrorCode,
				Message: err.Error(),
			})
		}
	}

	resp, err := ctrl.OrderService.GetOrdersByPagination(ctx, page, rowOfPage)
	if err != nil {
		slog.Error("GetOrdersByPagination error:", err)
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
