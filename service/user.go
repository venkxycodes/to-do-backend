package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"to-do/contract"
	"to-do/domain"
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
	//err := u.cache.GetUserByUserName(ctx, user.Username)
	//if err != nil {
	//	log.Fatal("err-user-already-exists", err)
	//	return err
	//}
	createErr := u.userRepo.AddNewUser(ctx, &domain.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	})
	return createErr
}
