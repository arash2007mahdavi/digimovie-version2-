package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateMobileNumber(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	res, _:= regexp.MatchString(`^09([0-9]{9})$`, value)
	return res
}