package main

import (
	"github.com/omidhaqi/clean-web-api/api"
	"github.com/omidhaqi/clean-web-api/config"
	"github.com/omidhaqi/clean-web-api/data/cache"
	"github.com/omidhaqi/clean-web-api/data/db"
	"github.com/omidhaqi/clean-web-api/data/db/migrations"
	"github.com/omidhaqi/clean-web-api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
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
	migrations.Up_1()

	api.InitServer(cfg)
}
