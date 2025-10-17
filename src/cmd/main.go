package main

import (
	"log"

	"github.com/mhmojtaba/golang-car-web-api/api"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/cache"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/migration"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name authorization
func main() {
	cfg := config.GetConfig()
	log.Printf("Config loaded: %+v", cfg)

	logger := logging.NewLogger(cfg)
	log.Println("Logger initialized")

	err := cache.InitRedis(cfg)
	if err != nil {
		log.Printf("Failed to initialize Redis: %v", err)
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()
	log.Println("Redis initialized")

	err = db.InitDb(cfg)
	if err != nil {
		log.Printf("Failed to initialize DB: %v", err)
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()
	log.Println("Database initialized")

	migration.Up_1()
	log.Println("Migration completed")

	api.InitServer(cfg)
	log.Println("Server initialized")
}
