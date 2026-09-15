package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"to-do/auth"
	"to-do/contract"
	"to-do/service"
	"to-do/utils"
)

type ToDoHandler struct {
	taskService service.TaskService
}

func NewToDoHandler(toDoService service.TaskService) ToDoHandler {
	return ToDoHandler{taskService: toDoService}
}

func (t *ToDoHandler) CreateTask(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var createTaskRequest contract.CreateTask
	if err := c.ShouldBindBodyWithJSON(&createTaskRequest); err != nil {
		log.Println(err.Error())
		httpStatus, errResp := utils.RenderError(errors.ErrUnsupported, createTaskRequest.Validate(), "Invalid request body")
		c.JSON(httpStatus, errResp)
		return
	}
	claims, ok := auth.ClaimsFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	createTaskRequest.UserName = claims.Username
	createTaskRequest.CreatedBy = claims.Username
	err := t.taskService.CreateTask(c, &createTaskRequest)
	if err != nil {
		log.Print(err)
		httpStatus, errorMessage := utils.RenderError(err, "Failed to create taskService")
		c.JSON(httpStatus, errorMessage)
		return
	}
	c.JSON(http.StatusCreated, utils.RenderSuccess("Task created successfully"))
	return
}

func (t *ToDoHandler) UpdateTask(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var updateTaskRequest contract.UpdateTask
	if err := c.ShouldBindBodyWithJSON(&updateTaskRequest); err != nil {
		log.Println(err.Error())
		httpStatus, errResp := utils.RenderError(errors.ErrUnsupported, updateTaskRequest.Validate(), "Invalid request body")
		c.JSON(httpStatus, errResp)
		return
	}
	claims, ok := auth.ClaimsFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	updateTaskRequest.UserName = claims.Username
	updateTaskRequest.UpdatedBy = claims.Username
	err := t.taskService.UpdateTask(c, &updateTaskRequest)
	if err != nil {
		log.Print(err.Error())
		httpStatus, errorMessage := utils.RenderError(err, "Failed to update taskService")
		c.JSON(httpStatus, errorMessage)
		return
	}
	c.JSON(http.StatusOK, utils.RenderSuccess("Task updated successfully"))
	return
}

func (t *ToDoHandler) GetTasks(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	username := c.Query("user_name")
	claims, ok := auth.ClaimsFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	if username != "" && username != claims.Username {
		c.JSON(http.StatusForbidden, gin.H{"error": "cannot access another user's tasks"})
		return
	}
	username = claims.Username
	tasks, err := t.taskService.GetTasks(c, username)
	if err != nil {
		log.Print(err.Error())
		httpStatus, errorMessage := utils.RenderError(err, "Failed to get tasks")
		c.JSON(httpStatus, errorMessage)
		return
	}
	c.JSON(http.StatusOK, utils.RenderSuccess(tasks))
	return
}

func (t *ToDoHandler) UpdateTaskStatus(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var updateTaskStatusRequest contract.UpdateTaskStatus
	if err := c.ShouldBindBodyWithJSON(&updateTaskStatusRequest); err != nil {
		log.Println(err.Error())
		httpStatus, errResp := utils.RenderError(errors.ErrUnsupported, updateTaskStatusRequest.Validate(), "Invalid request body")
		c.JSON(httpStatus, errResp)
		return
	}
	claims, ok := auth.ClaimsFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	updateTaskStatusRequest.UserName = claims.Username
	updateTaskStatusRequest.UpdatedBy = claims.Username
	err := t.taskService.UpdateTaskStatus(c, &updateTaskStatusRequest)
	if err != nil {
		log.Print(err.Error())
		httpStatus, errorMessage := utils.RenderError(err, "Failed to update status")
		c.JSON(httpStatus, errorMessage)
		return
	}
	c.JSON(http.StatusOK, utils.RenderSuccess("Task updated successfully"))
	return
}
