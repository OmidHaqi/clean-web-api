package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/omidhaqi/clean-web-api/api/handlers"
	"github.com/omidhaqi/clean-web-api/config"
)

func User(routers *gin.RouterGroup,cfg *config.Config)  {
	var h = handlers.NewUserHandler(cfg)
	
	routers.POST("/send-otp", h.SendOtp)
	
}