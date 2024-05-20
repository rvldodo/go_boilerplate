package config

import (
	"github.com/spf13/viper"

	"github.com/rvldodo/boilerplate/lib/log"
)

type Config struct {
	Addrs          string
	AppName        string
	DBUser         string
	DBPass         string
	DBAddrs        string
	DBName         string
	SecretTokenJWT string
	JWTExpiredTime int64
	// add new config when needed
}

var Envs = initConfig()

func initConfig() Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Viper config error: %v", err)
	}

	return Config{
		Addrs:          viper.GetString("ADDRS"),
		AppName:        viper.GetString("APP_NAME"),
		DBUser:         viper.GetString("MYSQL_USER"),
		DBPass:         viper.GetString("MYSQL_PASSWORD"),
		DBAddrs:        viper.GetString("MYSQL_ADDRESS"),
		DBName:         viper.GetString("MYSQL_DATABASE"),
		SecretTokenJWT: viper.GetString("JWT_SECRET"),
		JWTExpiredTime: viper.GetInt64("JWT_EXPIRED"),
	}
}