package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"to-do/contract"
	"to-do/domain"
	"to-do/repo"
)

type toDoService struct {
	toDoRepo repo.ToDoRepository
}

type ToDoService interface {
	CreateTask(ctx *gin.Context, task *contract.CreateTask) error
	UpdateTask(ctx *gin.Context, task *contract.UpdateTask) error
}

func NewToDoService(toDoRepo repo.ToDoRepository) ToDoService {
	return &toDoService{
		toDoRepo: toDoRepo,
	}
}

func (t *toDoService) CreateTask(ctx *gin.Context, task *contract.CreateTask) error {
	err := t.toDoRepo.AddTask(ctx, domain.Task{
		Id:           primitive.NewObjectID(),
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
	return err
}

func (t *toDoService) UpdateTask(ctx *gin.Context, task *contract.UpdateTask) error {
	err := t.toDoRepo.EditTask(ctx, domain.Task{
		Id:       task.Id,
		Name:     task.Name,
		Deadline: task.Deadline,
		Priority: task.Priority,
		Notes:    task.Notes,
		UpsertMeta: domain.UpsertMeta{
			UpdatedAt: time.Now().UnixMilli(),
			UpdatedBy: task.UpdatedBy,
		},
	})
	return err
}
