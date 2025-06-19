package database

import (
	"context"
	"fmt"
	"time"

	"github/Babe-piya/order-management/appconfig"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(config appconfig.Database) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.Hostname, config.Port, config.DatabaseName, config.Username, config.Password)
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = config.MaxPoolConnection
	poolConfig.MinConns = config.MinPoolConnection
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = time.Minute * 30
	poolConfig.HealthCheckPeriod = time.Minute * 1
	poolConfig.ConnConfig.ConnectTimeout = time.Second * 5
	poolConfig.ConnConfig.RuntimeParams = map[string]string{
		"timezone": config.Timezone,
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
