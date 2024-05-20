package config

import (
	"github.com/spf13/viper"

	"github.com/rvldodo/boilerplate/lib/log"
)

type Config struct {
	Addrs   string
	AppName string
	// add new config when needed
}

var Envs = initConfig()

func initConfig() Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Viper config error: %v", err)
	}

	return Config{
		Addrs:   viper.GetString("ADDRS"),
		AppName: viper.GetString("APP_NAME"),
	}
}
