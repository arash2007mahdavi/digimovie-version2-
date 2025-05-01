package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatePassword(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	res, _:= regexp.MatchString(`^[@0-9a-zA-Z]{5,50}$`, value)
	return res
}