package appconfig

type AppConfig struct {
	ServerPort string
	Database   Database
}

type Database struct {
	Hostname     string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}
