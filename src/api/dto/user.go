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

type TokenDetails struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}

type RegisterUserByUsernameRequest struct {
	FirstName string `json:"firstName" binding:"required,min=3"`
	LastName  string `json:"lastName" binding:"required,min=6"`
	Username  string `json:"username" binding:"required,min=5"`
	Email     string `json:"email" binding:"min=6,email"`
	Password  string `json:"password" binding:"required,password,min=6"`
}

type RegisterLoginByMobileRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
	Otp          string `json:"otp" binding:"required,min=4,max=4"`
}

type LoginByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}
type LoginByMobileRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
	Otp          string `json:"otp" binding:"required,min=4,max=4"`
}
