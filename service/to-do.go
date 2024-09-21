package service

import (
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
	CreateTask(task *contract.CreateTask) error
}

func NewToDoService(toDoRepo repo.ToDoRepository) ToDoService {
	return &toDoService{
		toDoRepo: toDoRepo,
	}
}

func (t *toDoService) CreateTask(task *contract.CreateTask) error {
	err := t.toDoRepo.AddTask(domain.Task{
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
