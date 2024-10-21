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

type ToDoHandler struct {
	toDoService service.ToDoService
}

func NewToDoHandler(toDoService service.ToDoService) ToDoHandler {
	return ToDoHandler{toDoService: toDoService}
}

func (t *ToDoHandler) CreateTask(c *gin.Context) {
	var createTaskRequest contract.CreateTask
	if err := c.ShouldBindBodyWithJSON(&createTaskRequest); err != nil {
		httpStatus, errResponse := utils.RenderError(errors.ErrUnsupported, createTaskRequest.Validate(), "Invalid request body")
		c.JSON(httpStatus, errResponse)
		return
	}
	err := t.toDoService.CreateTask(c, &createTaskRequest)
	if err != nil {
		log.Print(err)
		httpStatus, errorMessage := utils.RenderError(err, "Failed to create task")
		c.JSON(httpStatus, errorMessage)
		return
	}
	c.JSON(http.StatusCreated, utils.RenderSuccess("Task created successfully"))
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
		log.Print(err)
		httpStatus, errorMessage := utils.RenderError(err, "Failed to update task")
		c.JSON(httpStatus, errorMessage)
		return
	}
	c.JSON(http.StatusOK, utils.RenderSuccess("Task updated successfully"))
	return
}
