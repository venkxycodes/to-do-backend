package service

import (
	"github.com/gin-gonic/gin"
	"to-do/contract"
	"to-do/repo"
)

type userService struct {
	userRepo repo.UserRepository
}

type UserService interface {
	CreateUser(ctx *gin.Context, task *contract.CreateUser) error
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) CreateUser(ctx *gin.Context, user *contract.CreateUser) error {
	_, err := u.userRepo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	return nil
}
