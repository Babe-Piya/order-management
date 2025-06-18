package database

import (
	"context"
	"fmt"

	"github/Babe-piya/order-management/appconfig"

	"github.com/jackc/pgx/v5"
)

// TODO: use pool
func NewConnection(config appconfig.Database) (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Username, config.Password, config.Hostname,
		config.Port, config.DatabaseName)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	//defer conn.Close(context.Background())

	return conn, nil
}
