package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}
// UserByID godoc
// @Summary UserByID 
// @Description Check the server status
// @Tags Test
// @Accept json
// @Produce json
// @Param id path int true "User ID" 
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v2/test/user/{id} [post]
func (t *TestHandler) TestHandler(c *gin.Context) {
	id:=c.Param("id")
	c.JSON(http.StatusOK,gin.H{"message": "Hello from TestHandler with id: " + id})

}
