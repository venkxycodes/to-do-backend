package contract

import (
	"reflect"
	"testing"
)

func TestCreateUser_Validate(t *testing.T) {
	type fields struct {
		Name        string
		Username    string
		Password    string
		PhoneNumber string
	}
	tests := []struct {
		name string
		c    *SignUpUser
		want map[string]string
	}{
		{
			name: "test valid request",
			c: &SignUpUser{
				Name:        "venkat raman kannan",
				Username:    "venkatramankannanxo",
				Password:    "ThisIsAGoodPassword@99On100",
				PhoneNumber: "9900903821",
			},
			want: map[string]string{},
		},
		{
			name: "test invalid request",
			c: &SignUpUser{
				Name:        "",
				Username:    "1920234",
				Password:    "thisnot",
				PhoneNumber: "23900921",
			},
			want: map[string]string{
				"name":         "err-name-is-required",
				"username":     "err-username-should-not-be-lesser-than-8-characters",
				"password":     "err-password-must-be-atleast-8-characters-long",
				"phone_number": "err-phone-number-is-invalid",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{
			name:     "test valid password",
			password: "ThisIsAGoodPassword@99On100",
			want:     "",
		},
		{
			name:     "test less characters",
			password: "less",
			want:     "err-password-must-be-atleast-8-characters-long",
		},
		{
			name:     "test no digit",
			password: "lesscharacterpassword",
			want:     "err-password-must-contain-atleast-1-digit",
		},
		{
			name:     "test no uppercase letter",
			password: "nouppercase123",
			want:     "err-password-must-contain-atleast-1-uppercase-letter",
		},
		{
			name:     "test no special character",
			password: "NoUpperCase123",
			want:     "err-password-must-contain-atleast-1-special-character",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePassword(tt.password); got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
