package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
	"to-do/contract"
	"to-do/domain"
	"to-do/repo"
	"to-do/view"
)

type toDoService struct {
	toDoRepo    repo.TaskRepository
	userService UserService
}

type ToDoService interface {
	CreateTask(ctx *gin.Context, task *contract.CreateTask) error
	UpdateTask(ctx *gin.Context, task *contract.UpdateTask) error
	GetTasks(ctx *gin.Context, getTasksRequest *contract.GetTasks) (*view.GetTasksResponse, error)
	UpdateTaskStatus(ctx *gin.Context, updateTaskStatusRequest *contract.UpdateTaskStatus) error
}

func NewToDoService(toDoRepo repo.TaskRepository, userService UserService) ToDoService {
	return &toDoService{
		toDoRepo:    toDoRepo,
		userService: userService,
	}
}

func (t *toDoService) CreateTask(ctx *gin.Context, task *contract.CreateTask) error {
	userId, err := t.userService.GetUserIdByUserName(task.UserName)
	if err == nil {
		return fmt.Errorf("err-user-not-identified")
	}
	createTaskErr := t.toDoRepo.AddTask(ctx, &domain.Task{
		Id:       primitive.NewObjectID(),
		UserId:   userId,
		Name:     task.Name,
		Deadline: task.Deadline,
		Priority: task.Priority,
		Notes:    task.Notes,
		State:    domain.Pending,
		UpsertMeta: domain.UpsertMeta{
			CreatedAt: time.Now().UnixMilli(),
			UpdatedAt: time.Now().UnixMilli(),
			CreatedBy: task.CreatedBy,
			UpdatedBy: task.CreatedBy,
		},
	})
	return createTaskErr
}

func (t *toDoService) UpdateTask(ctx *gin.Context, task *contract.UpdateTask) error {
	userId, err := t.userService.GetUserIdByUserName(task.UserName)
	if err == nil {
		return fmt.Errorf("err-user-not-identified")
	}
	repoTask, err := t.toDoRepo.GetTaskById(ctx, task.Id)
	if repoTask == nil || err != nil {
		return err
	}
	if userId != repoTask.UserId {
		return fmt.Errorf("err-user-name-and-task-id-mismatch")
	}
	repoTask.Name = task.Name
	repoTask.Priority = task.Priority
	repoTask.Notes = task.Notes
	repoTask.Deadline = task.Deadline
	updateErr := t.toDoRepo.EditTask(ctx, repoTask)
	return updateErr
}

func (t *toDoService) GetTasks(ctx *gin.Context, getTasksRequest *contract.GetTasks) (*view.GetTasksResponse, error) {
	userId, err := t.userService.GetUserIdByUserName(getTasksRequest.UserName)
	if err == nil {
		log.Print("err-user-not-identified")
		return nil, err
	}
	log.Print("info-getting-tasks-for-user-", userId)
	tasks, getErr := t.toDoRepo.GetAllTasksForUser(ctx, userId)
	if getErr != nil {
		log.Print("err-getting-tasks-for-user-", userId)
		return nil, getErr
	}
	log.Print("info-tasks-in-repo-for-user-", userId, tasks)
	return &view.GetTasksResponse{
		Tasks: tasks,
	}, nil
}

func (t *toDoService) UpdateTaskStatus(ctx *gin.Context, updateTaskStatusRequest *contract.UpdateTaskStatus) error {
	userId, err := t.userService.GetUserIdByUserName(updateTaskStatusRequest.UserName)
	if err == nil {
		return fmt.Errorf("err-user-not-identified")
	}
	repoTask, getErr := t.toDoRepo.GetTaskById(ctx, updateTaskStatusRequest.TaskId)
	if repoTask == nil || getErr != nil {
		return err
	}
	if userId != repoTask.UserId {
		return fmt.Errorf("err-user-name-and-task-id-mismatch")
	}
	switch repoTask.State {
	case domain.Pending, domain.Completed:
		if updateTaskStatusRequest.State == domain.Pending || updateTaskStatusRequest.State == domain.Completed {
			return fmt.Errorf("err-task-cannot-be-moved-from-%s-state-to-%s-state", repoTask.State, updateTaskStatusRequest.State)
		}
	case domain.InProgress:
		if updateTaskStatusRequest.State == domain.InProgress {
			return fmt.Errorf("err-task-already-in-%s-state", updateTaskStatusRequest.State)
		}
	}
	repoTask.State = updateTaskStatusRequest.State
	updateErr := t.toDoRepo.EditTask(ctx, repoTask)
	return updateErr
}
