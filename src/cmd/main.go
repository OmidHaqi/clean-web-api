package main

import (
	"log"

	api "github.com/omidhaqi/clean-web-api/api"

	"github.com/omidhaqi/clean-web-api/config"

	"github.com/omidhaqi/clean-web-api/infra/cache"

	database "github.com/omidhaqi/clean-web-api/infra/persistence/database"

	"github.com/omidhaqi/clean-web-api/pkg/logging"
)

// @SecurityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	api.InitServer()

}
