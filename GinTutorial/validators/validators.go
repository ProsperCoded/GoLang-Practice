package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)
func IsCool(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().Elem().String(), "cool")
}