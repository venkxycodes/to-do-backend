package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var ErrInvalidToken = errors.New("err-invalid-token")

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Expires  int64  `json:"expires"`
}

func IssueToken(secret string, claims Claims) (string, error) {
	if secret == "" || claims.Username == "" || claims.UserID == 0 {
		return "", ErrInvalidToken
	}
	claims.Expires = time.Now().Add(24 * time.Hour).Unix()
	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString(payload)
	signature := sign(secret, encodedPayload)
	return encodedPayload + "." + signature, nil
}

func VerifyToken(secret, token string) (Claims, error) {
	parts := strings.Split(token, ".")
	if secret == "" || len(parts) != 2 || !hmac.Equal([]byte(sign(secret, parts[0])), []byte(parts[1])) {
		return Claims{}, ErrInvalidToken
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return Claims{}, ErrInvalidToken
	}
	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil || claims.Username == "" || claims.UserID == 0 || claims.Expires <= time.Now().Unix() {
		return Claims{}, ErrInvalidToken
	}
	return claims, nil
}

func sign(secret, payload string) string {
	h := hmac.New(sha256.New, []byte(secret))
	_, _ = h.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
