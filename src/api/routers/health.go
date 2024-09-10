package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/omidhaqi/clean-web-api/api/handlers"
)

func Health (r *gin.RouterGroup){

	handler := handlers.NewHealthHandler()
	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
	r.POST("/:id", handler.HealthById)
	
}