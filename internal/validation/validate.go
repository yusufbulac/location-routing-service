package validation

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = validator.New()
	_ = Validator.RegisterValidation("hexcolor", func(fl validator.FieldLevel) bool {
		str, ok := fl.Field().Interface().(string)
		if !ok {
			return false
		}
		return IsHexColor(str)
	})
}
