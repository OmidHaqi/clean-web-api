package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/omidhaqi/clean-web-api/api/handlers"
	"github.com/omidhaqi/clean-web-api/api/middlewares"
	"github.com/omidhaqi/clean-web-api/config"
)

func User(routers *gin.RouterGroup,cfg *config.Config)  {
	var h = handlers.NewUserHandler(cfg)
	
	routers.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	routers.POST("/login-by-username", h.LoginByUsername)
	routers.POST("/register-by-username", h.RegisterByUsername)
	routers.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}