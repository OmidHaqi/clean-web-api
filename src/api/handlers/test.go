package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) TestHandler(c *gin.Context) {

	c.JSON(http.StatusOK, "Test is UP")

}
