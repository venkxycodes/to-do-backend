package utils

import (
	"regexp"
)

// ValidatePhoneNumber validates if the given phone number is a valid Indian number (10 digits).
func ValidatePhoneNumber(phone string) bool {
	// Regular expression to match exactly 10 digits.
	var phoneRegex = regexp.MustCompile(`^[6-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}
