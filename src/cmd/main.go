package main

import (
	"log"

	"github.com/mhmojtaba/golang-car-web-api/api"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/cache"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
