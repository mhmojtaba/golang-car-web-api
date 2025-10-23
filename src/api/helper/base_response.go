package helper

import "github.com/mhmojtaba/golang-car-web-api/api/validation"

type BaseHttpResponse struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       int                           `json:"resultCode"`
	ValidationErrors *[]validation.ValidationError `json:"validationErrors"`
	Error            any                           `json:"Error"`
	Message          string                        `json:"message"`
}

func GenerateBaseResponse(result any, success bool, code int, msg string) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: code,
		Message:    msg,
	}
}

func GenerateBaseResponseWithError(result any, success bool, code int, err error, msg string) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: code,
		Error:      err.Error(),
		Message:    msg,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, code int, err error, msg string) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       code,
		ValidationErrors: validation.GetValidationErrors(err),
		Message:          msg,
	}
}
