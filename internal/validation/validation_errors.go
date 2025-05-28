package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	var sb strings.Builder
	for _, err := range err.(validator.ValidationErrors) {
		sb.WriteString(err.Field() + ": " + err.ActualTag() + "; ")
	}
	return sb.String()
}
