package main

import (
	"os"

	"github/Babe-piya/order-management/config"
	"github/Babe-piya/order-management/server"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	cfg := config.LoadFileConfig(configPath)
	e, db := server.Start(cfg)

	server.Shutdown(e, db)
}
