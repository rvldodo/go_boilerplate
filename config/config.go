package config

import (
	"github.com/spf13/viper"

	"github.com/rvldodo/boilerplate/lib/log"
)

type Config struct {
	Addrs              string
	AppName            string
	DBUser             string
	DBPass             string
	DBAddrs            string
	DBName             string
	SecretTokenJWT     string
	JWTExpiredTime     int64
	DBMigrations       string
	DBSeeds            string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRandomState  string
	RedisAddress       string
	RedisPassword      string
	RedisDB            int
	RedisTimeout       int64
	// add new config when needed
}

var Envs = initConfig()

func initConfig() Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Viper config error: %v", err)
	}

	return Config{
		Addrs:              viper.GetString("ADDRS"),
		AppName:            viper.GetString("APP_NAME"),
		DBUser:             viper.GetString("MYSQL_USER"),
		DBPass:             viper.GetString("MYSQL_PASSWORD"),
		DBAddrs:            viper.GetString("MYSQL_ADDRESS"),
		DBName:             viper.GetString("MYSQL_DATABASE"),
		SecretTokenJWT:     viper.GetString("JWT_SECRET"),
		JWTExpiredTime:     viper.GetInt64("JWT_EXPIRED"),
		DBMigrations:       viper.GetString("DB_MIGRATIONS"),
		DBSeeds:            viper.GetString("DB_SEEDS"),
		GoogleClientID:     viper.GetString("GOOGLE_CLIENTID"),
		GoogleClientSecret: viper.GetString("GOOGLE_CLIENTSECRET"),
		GoogleRandomState:  viper.GetString("GOOGLE_RANDOMSTATE"),
		RedisAddress:       viper.GetString("REDIS_ADDRESS"),
		RedisPassword:      viper.GetString("REDIS_PASSWORD"),
		RedisDB:            viper.GetInt("REDIS_DB"),
		RedisTimeout:       viper.GetInt64("REDIS_TIMEOUT"),
	}
}
