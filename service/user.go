package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"to-do/contract"
	"to-do/domain"
	"to-do/repo"
)

type userService struct {
	userRepo            repo.UserRepository
	usernameToUserIdMap *domain.UsernameToUserIdMap
}

type UserService interface {
	CreateUser(ctx *gin.Context, task *contract.SignUpUser) error
}

func NewUserService(userRepo repo.UserRepository) UserService {
	u := &userService{
		userRepo: userRepo,
	}
	u.usernameToUserIdMap = &domain.UsernameToUserIdMap{M: make(map[string]int64)}
	// Check if there are existing users on Db, if yes, populate them on the map
	// Existing in memory map gets cleared when we restart server
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	users, err := userRepo.GetAllUsers(ctx)
	if err != nil || len(users) == 0 {
		return u
	}
	for _, user := range users {
		u.usernameToUserIdMap.Set(user.Username, user.UserId)
	}
	return u
}

func (u *userService) CreateUser(ctx *gin.Context, user *contract.SignUpUser) error {
	userId, lastUserId := u.usernameToUserIdMap.Get(user.Username)
	fmt.Println(userId, lastUserId)
	if userId != 0 {
		return fmt.Errorf("err-username-already-exists")
	}
	createErr := u.userRepo.AddNewUser(ctx, &domain.User{
		Id:          primitive.NewObjectID(),
		Name:        user.Name,
		Username:    user.Username,
		Password:    user.Password,
		UserId:      lastUserId + 1,
		PhoneNumber: user.PhoneNumber,
	})
	u.usernameToUserIdMap.Set(user.Username, lastUserId+1)
	return createErr
}
