package utils

import "testing"

func TestValidatePhoneNumber(t *testing.T) {
	tests := []struct {
		name  string
		phone string
		want  bool
	}{
		{
			name:  "test indian phone number",
			phone: "9876543210",
			want:  true,
		},
		{
			name:  "test indian phone number 2",
			phone: "6789012345",
			want:  true,
		},
		{
			name:  "test empty string",
			phone: "",
			want:  false,
		},
		{
			name:  "test invalid phone number",
			phone: "1234567890",
			want:  false,
		},
		{
			name:  "test long phone number",
			phone: "98765432101",

			want: false,
		},
		{
			name:  "test number with texts",
			phone: "98765abcde",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePhoneNumber(tt.phone); got != tt.want {
				t.Errorf("ValidatePhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
