package appconfig

import "time"

type AppConfig struct {
	ServerPort string
	Database   Database
	Timeout    time.Duration
}

type Database struct {
	Hostname          string
	Port              string
	Username          string
	Password          string
	DatabaseName      string
	MaxPoolConnection int32
	MinPoolConnection int32
	Timezone          string
}
