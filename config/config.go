package config

import (
	"log"
	"path/filepath"
	"strings"

	"github/Babe-piya/order-management/appconfig"

	"github.com/spf13/viper"
)

func LoadFileConfig(configPath string) *appconfig.AppConfig {
	if len(configPath) == 0 {
		configPath = "env/config.yaml"
	}

	dir := filepath.Dir(configPath)
	filebase := filepath.Base(configPath)
	// file name without extension
	filename := strings.TrimSuffix(filebase, filepath.Ext(filebase))

	viper.SetConfigName(filename)
	viper.AddConfigPath(dir)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	// convert _ to dot in env variable
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %+v", err)
	}

	var cfg appconfig.AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshaling config: %+v", err)
	}

	return &cfg
}
