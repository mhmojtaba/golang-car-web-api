package helper

import (
	"net/http"

	"github.com/mhmojtaba/golang-car-web-api/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{

	// OTP
	service_errors.OtpExists:  409,
	service_errors.OtpUsed:    409,
	service_errors.InvalidOtp: 400,

	// User
	service_errors.EmailExists:      409,
	service_errors.UsernameExists:   409,
	service_errors.UserNotFound:     404,
	service_errors.PermissionDenied: 403,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
