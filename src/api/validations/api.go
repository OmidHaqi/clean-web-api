package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/omidhaqi/clean-web-api/api/routers"
	"github.com/omidhaqi/clean-web-api/config"
)

func InitServer() {

	cfg := config.GetConfig() //1

	r := gin.New() //2

	r.Use(gin.Logger(), gin.Recovery()) //3

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{

		health := v1.Group("/health")
		routers.Health(health)
		healthById := v1.Group("/health-by-id")
		routers.Health(healthById)

		v2 := api.Group("/v2")
		{
			test := v2.Group("/test")
			routers.Test(test)
		}

		r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))

	}
}

// InitServer initializes and starts the web server for the clean-web-api application.
// It sets up the Gin framework, registers routes, and starts listening on the specified port.
//
// The function follows these steps:
// 1. Retrieves the configuration settings from the config file using config.GetConfig().
// 2. Creates a new Gin engine using gin.New().
// 3. Applies middleware for logging and recovery using r.Use(gin.Logger(), gin.Recovery()).
// .
// .
// .
// 5. Registers routes for health checks under "/health" and "/health-by-id" using routers.Health().
// 6. Starts the server to listen on the specified internal port using r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort)).
