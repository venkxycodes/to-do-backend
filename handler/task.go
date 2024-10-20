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
		httpStatus, errResponse := utils.RenderError(errors.ErrUnsupported, createUserRequest.Validate(), "Invalid request body")
		c.JSON(httpStatus, errResponse)
		return
	}
	err := t.toDoService.CreateTask(c, &createUserRequest)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (t *ToDoHandler) UpdateTask(c *gin.Context) {
	var updateTaskRequest contract.UpdateTask
	if err := c.ShouldBindBodyWithJSON(&updateTaskRequest); err != nil {
		httpStatus, errResponse := utils.RenderError(errors.ErrUnsupported, nil, "Invalid request body")
		c.JSON(httpStatus, errResponse)
		return
	}
	err := t.toDoService.UpdateTask(c, &updateTaskRequest)
	if err != nil {
		log.Fatal(err)
	}
	return
}
