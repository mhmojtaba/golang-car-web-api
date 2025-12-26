package service_errors

const (
	OtpExists        = "Otp Already Exists"
	OtpUsed          = "Otp Already Used"
	InvalidOtp       = "Invalid Otp"
	EmailExists      = "email is already exists"
	UsernameExists   = "username is already exists"
	RecordNotFound   = "user is not found"
	PermissionDenied = "permission is denied"

	UnExpectedError = "unexpected error occurred"
	ClaimsNotFound  = "claims not found in token"

	TokenRequired = "token is required"
	InvalidToken  = "invalid token"
	TokenExpired  = "token is expired"
	Forbidden     = "forbidden access"
)
