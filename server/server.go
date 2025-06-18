package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github/Babe-piya/order-management/appconfig"
	"github/Babe-piya/order-management/database"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func Start(config *appconfig.AppConfig) (*echo.Echo, *pgx.Conn) {
	db, err := database.NewConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	Routes(e, db, config)

	go func() {
		endPoint := fmt.Sprintf(":%s", config.ServerPort)
		if err := e.Start(endPoint); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Error(err.Error())
			e.Logger.Fatal("shutting down the server")
		}
	}()

	return e, db
}

func Shutdown(e *echo.Echo, db *pgx.Conn) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	defer func() {
		if err := db.Close(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()
}
