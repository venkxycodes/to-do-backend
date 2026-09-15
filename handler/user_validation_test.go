package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"to-do/contract"
)

type validationUserService struct {
	created bool
}

func (s *validationUserService) GetUserIdByUserName(string) (int64, error) { return 0, nil }
func (s *validationUserService) GetUserIdByUserNameWithContext(*gin.Context, string) (int64, error) {
	return 0, nil
}
func (s *validationUserService) CreateUser(*gin.Context, *contract.SignUpUser) error {
	s.created = true
	return nil
}
func (s *validationUserService) LoginUser(*gin.Context, *contract.LoginUser) error { return nil }
func (s *validationUserService) Authenticate(*gin.Context, *contract.LoginUser) (string, error) {
	return "", nil
}

func TestSignUpRejectsWeakPasswordAfterBinding(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := &validationUserService{}
	handler := NewUserHandler(service)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"Example","username":"example-user","password":"weak","phone_number":"9900923821"}`))
	context.Request.Header.Set("Content-Type", "application/json")

	handler.SignUpUser(context)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
	}
	if service.created {
		t.Fatal("CreateUser was called for invalid input")
	}
}
