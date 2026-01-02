package helper

import "github.com/mhmojtaba/golang-car-web-api/api/validation"

type BaseHttpResponse struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       ResultCode                    `json:"resultCode"`
	ValidationErrors *[]validation.ValidationError `json:"validationErrors"`
	Error            any                           `json:"Error"`
	Message          string                        `json:"message"`
}

func GenerateBaseResponse(result any, success bool, ResultCode ResultCode, msg string) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: ResultCode,
		Message:    msg,
	}
}

func GenerateBaseResponseWithError(result any, success bool, ResultCode ResultCode, err error, msg string) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: ResultCode,
		Error:      err.Error(),
		Message:    msg,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, ResultCode ResultCode, err error, msg string) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       ResultCode,
		ValidationErrors: validation.GetValidationErrors(err),
		Message:          msg,
	}
}
