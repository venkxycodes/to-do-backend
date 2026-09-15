package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"to-do/contract"
	appErrors "to-do/error"
	"to-do/service"
	"to-do/utils"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (u UserHandler) SignUpUser(c *gin.Context) {
	var createUserRequest contract.SignUpUser
	if err := c.ShouldBindBodyWithJSON(&createUserRequest); err != nil {
		httpStatus, errResponse := utils.RenderError(errors.ErrUnsupported, createUserRequest.Validate(), "Invalid request body")
		c.JSON(httpStatus, errResponse)
		return
	}
	if validationErrors := createUserRequest.Validate(); len(validationErrors) > 0 {
		httpStatus, errResponse := utils.RenderError(appErrors.ErrInvalidRequest, validationErrors, "Invalid request body")
		c.JSON(httpStatus, errResponse)
		return
	}
	err := u.userService.CreateUser(c, &createUserRequest)
	if err != nil {
		log.Print(err)
		httpStatus, errorMessage := utils.RenderError(err, "Failed to sign up user")
		c.JSON(httpStatus, errorMessage)
		return
	}
	c.JSON(http.StatusCreated, utils.RenderSuccess("User signed up successfully"))
	return
}

func (u UserHandler) LoginUser(c *gin.Context) {
	var loginUserRequest contract.LoginUser
	if err := c.ShouldBindBodyWithJSON(&loginUserRequest); err != nil {
		httpStatus, errResponse := utils.RenderError(errors.ErrUnsupported, loginUserRequest.Validate(), "Invalid request body")
		c.JSON(httpStatus, errResponse)
		return
	}
	if loginUserRequest.Username == "" || loginUserRequest.Password == "" {
		httpStatus, errResponse := utils.RenderError(appErrors.ErrInvalidRequest, "username and password are required", "Invalid request body")
		c.JSON(httpStatus, errResponse)
		return
	}
	token, err := u.userService.Authenticate(c, &loginUserRequest)
	if err != nil {
		log.Print(err)
		httpStatus, errorMessage := utils.RenderError(err, "Failed to login user")
		c.JSON(httpStatus, errorMessage)
		return
	}
	c.JSON(http.StatusOK, utils.RenderSuccess(map[string]string{"token": token}))
	return
}
