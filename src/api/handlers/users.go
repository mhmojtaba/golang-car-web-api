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

// login by username godoc
// @Summary Login By Username
// @Description Login By Username function
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.LoginByUsernameRequest true "Login By Username Request"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/users/login-by-username/ [post]
func (u *UsersHandler) LoginByUsername(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"),
		)
		return
	}
	tokenDetails, err := u.service.LoginByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err, err.Error()))

		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(tokenDetails, true, 0, "Login successful"))
}

// register by username godoc
// @Summary Register By Username
// @Description Register By Username function
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserByUsernameRequest true "Register By Username Request"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/users/register-by-username/ [post]
func (u *UsersHandler) RegisterByUsername(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"),
		)
		return
	}
	err = u.service.RegisterUserByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err, err.Error()))

		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0, "User registered successfully"))
}

// register by mobile number godoc
// @Summary Register By Mobile Number
// @Description Register By Mobile Number function
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.RegisterLoginByMobileRequest true "Register By Mobile Number Request"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/users/login-by-mobile/ [post]
func (u *UsersHandler) RegisterLoginByMobileNumber(c *gin.Context) {
	req := new(dto.RegisterLoginByMobileRequest)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"),
		)
		return
	}
	tokenDetails, err := u.service.RegisterLoginByMobileNumber(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err, err.Error()))

		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(tokenDetails, true, 0, "User registered and logged in successfully"))
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
			helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"),
		)
		return
	}
	err = u.service.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, -1, err, err.Error()))

		return
	}
	// call sms service to send otp
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(&dto.SendOtpResponse{
		Message: "OTP has been sent successfully",
	}, true, 0, "OTP has been sent successfully"))

}
