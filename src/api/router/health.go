package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omidhaqi/clean-web-api/api/handler"
)

func Health(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()

	r.GET("/", handler.Health)
}
