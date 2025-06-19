package controller

import (
	"log"
	"net/http"
	"strconv"

	"github/Babe-piya/order-management/appconstant"

	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) GetOrdersByPagination(c echo.Context) error {
	ctx := c.Request().Context()
	page := 1
	var err error
	pageStr := c.QueryParam("page")
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			log.Println(err)

			return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
				Code:    "0",
				Message: err.Error(),
			})
		}
	}

	rowOfPage := 50
	rowOfPageStr := c.QueryParam("row_of_page")
	if rowOfPageStr != "" {
		rowOfPage, err = strconv.Atoi(rowOfPageStr)
		if err != nil {
			log.Println(err)

			return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
				Code:    "0",
				Message: err.Error(),
			})
		}
	}

	resp, err := ctrl.OrderService.GetOrdersByPagination(ctx, page, rowOfPage)
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusInternalServerError, appconstant.ErrorResponse{
			Code:    "0",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
