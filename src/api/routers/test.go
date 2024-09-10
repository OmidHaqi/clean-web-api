package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/omidhaqi/clean-web-api/api/handlers"
)

func Test(r *gin.RouterGroup) {

	handler := handlers.NewTestHandler()

	r.GET("/", handler.TestHandler)

}
