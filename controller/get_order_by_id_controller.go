package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) GetOrderByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusBadRequest, err)
	}

	resp, err := ctrl.OrderService.GetOrderByID(ctx, int64(orderID))
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}
