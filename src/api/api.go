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

