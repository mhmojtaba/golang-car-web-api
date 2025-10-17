package service_errors

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error
}

func (e *ServiceError) Error() string {
	return e.Message
}
