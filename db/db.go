package db

import (
	mysqlid "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/rvldodo/boilerplate/config"
	"github.com/rvldodo/boilerplate/lib/log"
)

var DB *gorm.DB

func init() {
	db, err := newMySQLStorage(mysqlid.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPass,
		Addr:                 config.Envs.DBAddrs,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Errorf("DB init error: %v", err)
	}

	initStorage(db)
	DB = db
}

func newMySQLStorage(config mysqlid.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: config.FormatDSN()}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Errorf("Gorm error: %v", err)
		return nil, err
	}

	return db, nil
}

func initStorage(db *gorm.DB) {
	sql, _ := db.DB()
	if err := sql.Ping(); err != nil {
		log.Errorf("Ping error: %v", err)
	}

	log.Info("DB Connected Succesfully")
}
