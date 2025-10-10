package helper

import "github.com/mhmojtaba/golang-car-web-api/api/validation"

type BaseHttpResponse struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       int                           `json:"resultCode"`
	ValidationErrors *[]validation.ValidationError `json:"validationErrors"`
	Error            any                           `json:"Error"`
}

func GenerateBaseResponse(result any, success bool, code int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: code,
	}
}

func GenerateBaseResponseWithError(result any, success bool, code int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: code,
		Error:      err.Error(),
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, code int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       code,
		ValidationErrors: validation.GetValidationErrors(err),
	}
}
