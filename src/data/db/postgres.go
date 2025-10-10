package db

import (
	"fmt"
	"log"
	"time"

	"github.com/mhmojtaba/golang-car-web-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	var err error
	// create connection string
	dsn := fmt.Sprintf("host=%s port=%s User=%s	Password=%s	DbName=%s SSLMode=%s TimeZone=Asia/Tehran", cfg.Postgres.Host, cfg.Server.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := dbClient.DB()

	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("db connection established")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	sqlDB, _ := dbClient.DB()
	sqlDB.Close()
}
