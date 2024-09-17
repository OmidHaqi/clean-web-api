package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omidhaqi/clean-web-api/api/dto"
	"github.com/omidhaqi/clean-web-api/api/helper"
	"github.com/omidhaqi/clean-web-api/config"
	"github.com/omidhaqi/clean-web-api/services"
)

type UserHandler struct {
	services *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler{
	service := services.NewUserService(cfg)
	return &UserHandler{services: service}
}
// SendOtp godoc
// @Summary Send OTP to user
// @Description Send OTP to user 
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context)  {
	req := new(dto.GetOtpRequest)
	err:=c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,helper.GenerateBaseResponseWithValidationError(nil , false,-1,err))
		return
	}

	err = h.services.SendOtp(req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrToStatusCode(err),helper.GenerateBaseResponseWithError(nil,false,-1,err))
		return
	}
	//call internal SMS service
	c.JSON(http.StatusCreated,helper.GenerateBaseResponse(nil,true,-1))
}