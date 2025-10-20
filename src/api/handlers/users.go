package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/api/helper"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/services"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{
		service: service,
	}
}

// send otp godoc
// @Summary Send OTP
// @Description Send OTP function
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetOtpRequest true "Send OTP Request"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/users/send-otp/ [post]
func (u *UsersHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err),
		)
		return
	}
	err = u.service.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	// call sms service to send otp
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(&dto.SendOtpResponse{
		Message: "OTP has been sent successfully",
	}, true, 0))

}
