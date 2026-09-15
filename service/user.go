package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http/httptest"
	"to-do/auth"
	"to-do/config"
	"to-do/contract"
	"to-do/domain"
	"to-do/repo"
)

type userService struct {
	userRepo            repo.UserRepository
	usernameToUserIdMap *domain.UsernameToUserIdMap
}

type userIDAllocator interface {
	NextUserID(ctx *gin.Context) (int64, error)
}

type UserService interface {
	GetUserIdByUserName(username string) (int64, error)
	GetUserIdByUserNameWithContext(ctx *gin.Context, username string) (int64, error)
	CreateUser(ctx *gin.Context, user *contract.SignUpUser) error
	LoginUser(ctx *gin.Context, user *contract.LoginUser) error
	Authenticate(ctx *gin.Context, user *contract.LoginUser) (string, error)
}

func (u *userService) GetUserIdByUserNameWithContext(ctx *gin.Context, username string) (int64, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return 0, fmt.Errorf("err-user-not-identified")
	}
	return user.UserId, nil
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

// GetUserIdByUserName method returns userId and an error when a user already exists,
// returns the last userId and no error when a user does not exist
func (u *userService) GetUserIdByUserName(username string) (int64, error) {
	userId, lastUserId := u.usernameToUserIdMap.Get(username)
	if userId != 0 {
		return userId, fmt.Errorf("err-username-already-taken")
	}
	return lastUserId, nil
}

func (u *userService) CreateUser(ctx *gin.Context, user *contract.SignUpUser) error {
	userID := int64(0)
	if allocator, ok := u.userRepo.(userIDAllocator); ok {
		var err error
		userID, err = allocator.NextUserID(ctx)
		if err != nil {
			return err
		}
	} else {
		var err error
		lastUserId, err := u.GetUserIdByUserName(user.Username)
		if err != nil {
			return err
		}
		userID = lastUserId + 1
	}
	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return hashErr
	}
	createErr := u.userRepo.AddNewUser(ctx, &domain.User{
		Id:          primitive.NewObjectID(),
		Name:        user.Name,
		Username:    user.Username,
		Password:    string(passwordHash),
		UserId:      userID,
		PhoneNumber: user.PhoneNumber,
	})
	if createErr != nil {
		return createErr
	}
	u.usernameToUserIdMap.Set(user.Username, userID)
	return nil
}

func (u *userService) LoginUser(ctx *gin.Context, userLoginInfo *contract.LoginUser) error {
	userId, err := u.GetUserIdByUserName(userLoginInfo.Username)
	if err == nil {
		return fmt.Errorf("err-username-not-identified")
	}
	userDetails, getUserErr := u.userRepo.GetUserByUserId(ctx, userId)
	if getUserErr != nil {
		return getUserErr
	}
	if bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(userLoginInfo.Password)) != nil {
		return fmt.Errorf("err-incorrect-password")
	}
	return nil
}

func (u *userService) Authenticate(ctx *gin.Context, userLoginInfo *contract.LoginUser) (string, error) {
	if err := u.LoginUser(ctx, userLoginInfo); err != nil {
		return "", err
	}
	userID, err := u.GetUserIdByUserName(userLoginInfo.Username)
	if err == nil {
		return "", fmt.Errorf("err-username-not-identified")
	}
	return auth.IssueToken(config.GetConfig().JWTSecret, auth.Claims{UserID: userID, Username: userLoginInfo.Username})
}
