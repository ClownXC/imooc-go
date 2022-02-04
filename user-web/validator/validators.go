package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateMobile(f validator.FieldLevel) bool {

	mobile := f.Field().String()
	ok, _ := regexp.MatchString(`^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`, mobile)
	return ok
}
