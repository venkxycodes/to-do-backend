package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRequireAuthRejectsMissingAndInvalidTokens(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(RequireAuth("secret"))
	router.GET("/protected", func(c *gin.Context) { c.Status(http.StatusNoContent) })

	for _, token := range []string{"", "Bearer invalid"} {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/protected", nil)
		if token != "" {
			request.Header.Set("Authorization", token)
		}
		router.ServeHTTP(recorder, request)
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("token %q: status = %d, want %d", token, recorder.Code, http.StatusUnauthorized)
		}
	}
}
