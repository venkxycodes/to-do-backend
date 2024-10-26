package repo

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"to-do/domain"
)

type UserRepoMock struct {
	mock.Mock
}

func (u UserRepoMock) AddNewUser(ctx *gin.Context, user *domain.User) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

func (u UserRepoMock) GetUserByUsername(ctx *gin.Context, username string) (*domain.User, error) {
	args := u.Called(ctx, username)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (u UserRepoMock) GetAllUsers(ctx *gin.Context) ([]domain.User, error) {
	args := u.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}
