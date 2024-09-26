package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/omidhaqi/clean-web-api/api/middlewares"
	"github.com/omidhaqi/clean-web-api/api/routers"
	validation "github.com/omidhaqi/clean-web-api/api/validations"

	"github.com/omidhaqi/clean-web-api/config"
	"github.com/omidhaqi/clean-web-api/docs"
	"github.com/omidhaqi/clean-web-api/pkg/logging"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {

	r := gin.New()

	RegisterValidators()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler) /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterSwagger(r, cfg)

	RegisterRoutes(r, cfg)

	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}

}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	{

		health := v1.Group("/health")
		routers.Health(health)

		healthById := v1.Group("/health-by-id")
		routers.Health(healthById)

		users := v1.Group("/users")
		routers.User(users, cfg)

		v2 := api.Group("/v2")
		{
			test := v2.Group("/test")
			routers.Test(test)
		}

		r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))

	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "clean web api"
	docs.SwaggerInfo.Description = "a simple web api by Go and Gin framework"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.InternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
