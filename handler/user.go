package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"to-do/contract"
	"to-do/service"
	"to-do/utils"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (u UserHandler) SignUp(c *gin.Context) {
	var createUserRequest contract.SignUpUser
	if err := c.ShouldBindBodyWithJSON(&createUserRequest); err != nil {
		httpStatus, errResponse := utils.RenderError(errors.ErrUnsupported, createUserRequest.Validate(), "Invalid request body")
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
