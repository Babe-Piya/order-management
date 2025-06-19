package controller

import (
	"log"
	"net/http"
	"strconv"

	"github/Babe-piya/order-management/appconstant"
	"github/Babe-piya/order-management/service"

	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) UpdateStatusByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("order_id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
			Code:    "0",
			Message: err.Error(),
		})
	}

	req := service.UpdateStatusByIDRequest{}
	err = c.Bind(&req)
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusBadRequest, appconstant.ErrorResponse{
			Code:    "0",
			Message: err.Error(),
		})
	}
	req.ID = int64(orderID)
	resp, err := ctrl.OrderService.UpdateStatusByID(ctx, req)
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusInternalServerError, appconstant.ErrorResponse{
			Code:    "0",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
