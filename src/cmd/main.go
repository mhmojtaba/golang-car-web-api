package main

import (
	"github.com/mhmojtaba/golang-car-web-api/api"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/cache"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name authorization
func main() {
	cfg := config.GetConfig()

	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
}
