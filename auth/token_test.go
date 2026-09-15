package auth

import (
	"testing"
	"time"
)

func TestIssueAndVerifyToken(t *testing.T) {
	token, err := IssueToken("test-secret", Claims{UserID: 7, Username: "someone"})
	if err != nil {
		t.Fatalf("IssueToken() error = %v", err)
	}
	claims, err := VerifyToken("test-secret", token)
	if err != nil {
		t.Fatalf("VerifyToken() error = %v", err)
	}
	if claims.UserID != 7 || claims.Username != "someone" || claims.Expires <= time.Now().Unix() {
		t.Fatalf("unexpected claims: %+v", claims)
	}
}

func TestVerifyTokenRejectsTamperingAndWrongSecret(t *testing.T) {
	token, err := IssueToken("test-secret", Claims{UserID: 7, Username: "someone"})
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		name   string
		secret string
		token  string
	}{
		{name: "wrong secret", secret: "other-secret", token: token},
		{name: "tampered token", secret: "test-secret", token: token + "x"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := VerifyToken(test.secret, test.token); err == nil {
				t.Fatal("VerifyToken() unexpectedly accepted invalid token")
			}
		})
	}
}
