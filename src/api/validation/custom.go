package validation

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"propertyName"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, e := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = e.Field()
			el.Tag = e.Tag()
			el.Value = e.Param()

			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}
	return nil
}

// mobile regex for validation
const mobileRegex = `^09(1[0-9]|3[1-9]|2[1-9])-?[0-9]{3}-?[0-9]{4}$`

// email regex for validation
const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

// password regex for validation
const passwordRegex = `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`

func ValidateMobile(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	isMatched, err := regexp.MatchString(mobileRegex, value)
	if err != nil {
		return false
	}
	return isMatched
}

func ValidateEmail(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	isMatched, err := regexp.MatchString(emailRegex, value)
	if err != nil {
		return false
	}
	return isMatched
}

func ValidatePassword(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	isMatched, err := regexp.MatchString(passwordRegex, value)
	if err != nil {
		return false
	}
	return isMatched
}
