package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

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
