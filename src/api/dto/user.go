package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,min=11,max=11,mobile"`
}

type VerifyOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
	Otp          string `json:"otp" binding:"required"`
}

type SendOtpResponse struct {
	Message string `json:"message"`
}

type VerifyOtpResponse struct {
	Message string `json:"message"`
}
