package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler   {
	return &HealthHandler{} 
}

// HealthCheck godoc
// @Summary Health Check
// @Description Check the server status
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Working!")
	
}


func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "Working Post!")
	
}

func (h *HealthHandler) HealthById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK,fmt.Sprintf("Working PostById : %s!",id)) 
}
