package contract

import (
	"regexp"
	"to-do/utils"
)

type SignUpUser struct {
	Name        string `json:"name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (c *SignUpUser) Validate() map[string]string {
	errors := make(map[string]string)
	if c.Name == "" {
		errors["name"] = "err-name-is-required"
	}
	if len(c.Username) < 8 {
		errors["username"] = "err-username-should-not-be-lesser-than-8-characters"
	}
	if err := ValidatePassword(c.Password); err != "" {
		errors["password"] = err
	}
	if !utils.ValidatePhoneNumber(c.PhoneNumber) {
		errors["phone_number"] = "err-phone-number-is-invalid"
	}
	return errors
}

func ValidatePassword(password string) string {
	if len(password) < 8 {
		return "err-password-must-be-atleast-8-characters-long"
	}

	hasDigit := regexp.MustCompile(`[0-9]`)
	hasUpper := regexp.MustCompile(`[A-Z]`)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",<>\.\?\/\\|~]`)

	if !hasDigit.MatchString(password) {
		return "err-password-must-contain-atleast-1-digit"
	}
	if !hasUpper.MatchString(password) {
		return "err-password-must-contain-atleast-1-uppercase-letter"
	}
	if !hasSpecial.MatchString(password) {
		return "err-password-must-contain-atleast-1-special-character"
	}
	return ""
}
