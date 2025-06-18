package server

import (
	"net/http"

	"github/Babe-piya/order-management/appconfig"
	"github/Babe-piya/order-management/controller"
	"github/Babe-piya/order-management/repositories"
	"github/Babe-piya/order-management/service"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, db *pgx.Conn, config *appconfig.AppConfig) {
	orderRepo := repositories.NewOrderRepository(db)

	orderService := service.NewOrderService(orderRepo)

	orderCtrl := controller.NewOrderController(orderService)

	e.GET("/health", func(c echo.Context) error {
		response := map[string]string{
			"message": "service available",
		}
		return c.JSON(http.StatusOK, response)
	})

	orderAPI := e.Group("/orders")

	orderAPI.GET("/:id", orderCtrl.GetOrderByID)
}
