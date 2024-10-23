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
	toDoRepo    repo.ToDoRepository
	userService UserService
}

type ToDoService interface {
	CreateTask(ctx *gin.Context, task *contract.CreateTask) error
	UpdateTask(ctx *gin.Context, task *contract.UpdateTask) error
	GetTasks(ctx *gin.Context, task *contract.GetTasksRequest) (*view.GetTasksResponse, error)
}

func NewToDoService(toDoRepo repo.ToDoRepository, userService UserService) ToDoService {
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
		Id:           primitive.NewObjectID(),
		UserId:       userId,
		Name:         task.Name,
		Deadline:     task.Deadline,
		Priority:     task.Priority,
		Notes:        task.Notes,
		CurrentState: domain.Pending,
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
	repoTask, err := t.toDoRepo.GetTaskById(ctx, task.Id)
	if repoTask == nil || err != nil {
		return err
	}
	repoTask.Name = task.Name
	repoTask.Priority = task.Priority
	repoTask.Notes = task.Notes
	repoTask.Deadline = task.Deadline
	updateErr := t.toDoRepo.EditTask(ctx, repoTask)
	return updateErr
}

func (t *toDoService) GetTasks(ctx *gin.Context, task *contract.GetTasksRequest) (*view.GetTasksResponse, error) {
	userId, err := t.userService.GetUserIdByUserName(task.UserName)
	if err == nil {
		log.Print("err-user-not-identified")
		return nil, err
	}
	log.Print("info-getting-tasks-for-user-", userId)
	tasks, err := t.toDoRepo.GetAllTasksForUser(ctx, userId)
	if err != nil {
		log.Print("err-getting-tasks-for-user-", userId)
		return nil, err
	}
	log.Print(tasks)
	tasksResponse := &view.GetTasksResponse{}
	for _, task := range tasks {
		tasksResponse.Tasks = append(tasksResponse.Tasks, task)
	}
	return tasksResponse, nil
}
