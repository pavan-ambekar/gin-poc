package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func IsCool(title validator.FieldLevel) bool {
	return strings.Contains(title.Field().String(), "cool")
}
