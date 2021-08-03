package utils

import (
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// NewValidate 构造验证器
func NewValidate() *validator.Validate {
	validate = validator.New()
	validate.RegisterValidation("is_price", ValidatePrice)
	validate.RegisterValidation("is_username", ValidateUsername)
	validate.RegisterValidation("is_password", ValidatePassword)
	return validate
}

// ValidatePrice 自定义金额验校验
func ValidatePrice(fl validator.FieldLevel) bool {
	price := fl.Field().String()

	return regexp.MustCompile(`(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^(0){1}$)|(^[0-9]\.[0-9]([0-9])?$)`).MatchString(price)
}

// ValidateUsername 自定义用户名校验
func ValidateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	return regexp.MustCompile(`^[a-zA-Z0-9]{5,15}$`).MatchString(username)
}

// ValidatePassword 自定义密码校验
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if regexp.MustCompile(`\s`).MatchString(password) {
		return false
	}

	var (
		isUpper   = false
		isLower   = false
		isNumber  = false
		isSpecial = false
	)
	// 8-25 限制
	if len(password) < 8 || len(password) > 25 {
		return false
	}

	for _, s := range password {
		switch {
		case unicode.IsUpper(s):
			isUpper = true
		case unicode.IsLower(s):
			isLower = true
		case unicode.IsNumber(s):
			isNumber = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			isSpecial = true
		default:
		}
	}

	if (isUpper || isLower) && isNumber && isSpecial {
		return true
	}

	return false
}
