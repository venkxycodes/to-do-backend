package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"to-do/contract"
	"to-do/domain"
	"to-do/repo"
)

type userService struct {
	cache    repo.UserCache
	userRepo repo.UserRepository
}

type UserService interface {
	CreateUser(ctx *gin.Context, task *contract.CreateUser) error
}

func NewUserService(cache repo.UserCache, userRepo repo.UserRepository) UserService {
	return &userService{
		cache:    cache,
		userRepo: userRepo,
	}
}

func (u *userService) CreateUser(ctx *gin.Context, user *contract.CreateUser) error {
	_, err := u.userRepo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		log.Fatal("err-user-already-exists", err)
		return err
	}
	createErr := u.userRepo.AddNewUser(ctx, &domain.User{
		Id: primitive.NewObjectID(),

		Username: user.Username,
		Password: user.Password,
	})
	return createErr
}
