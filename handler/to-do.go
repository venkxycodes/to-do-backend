package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"to-do/contract"
	"to-do/service"
	"to-do/utils"
)

type ToDoHandler struct {
	toDoService service.ToDoService
}

func NewToDoHandler(toDoService service.ToDoService) ToDoHandler {
	return ToDoHandler{toDoService: toDoService}
}

func (t *ToDoHandler) CreateToDo(c *gin.Context) {
	var createUserRequest contract.CreateTask
	if err := c.ShouldBindBodyWithJSON(&createUserRequest); err != nil {
		httpStatus, errResponse := utils.RenderError(errors.ErrUnsupported, nil, "Invalid request body")
		c.JSON(httpStatus, errResponse)
		return
	}
	err := t.toDoService.CreateTask(&createUserRequest)
	if err != nil {
		log.Fatal(err)
	}
	return
}
