package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/omidhaqi/clean-web-api/api/middlewares"
	"github.com/omidhaqi/clean-web-api/api/routers"
	"github.com/omidhaqi/clean-web-api/config"
	"github.com/omidhaqi/clean-web-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer() {

	cfg := config.GetConfig()

	r := gin.New() 

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery())  


    RegisterSwagger(r, cfg)

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

func RegisterSwagger(r *gin.Engine , cfg *config.Config){
	docs.SwaggerInfo.Title = "clean web api"
	docs.SwaggerInfo.Description = "a simple web api by Go and Gin framework"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s",cfg.Server.InternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

