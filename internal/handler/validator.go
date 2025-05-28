package handler

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

func FormatValidationError(err error) string {
	var sb strings.Builder
	for _, err := range err.(validator.ValidationErrors) {
		sb.WriteString(err.Field() + ": " + err.ActualTag() + "; ")
	}
	return sb.String()
}
