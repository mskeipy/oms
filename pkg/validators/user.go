package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"unicode"
)

func ValidatePassword(fl validator.FieldLevel) bool {
	hasLetter, hasDigit, hasSpecial := false, false, false
	specialChars := "@$!%*?&"
	password := fl.Field().String()
	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsRune(specialChars, char):
			hasSpecial = true
		}
	}

	return hasLetter && hasDigit && hasSpecial
}
