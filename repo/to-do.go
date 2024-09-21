package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"to-do/domain"
)

type toDoRepo struct {
	dbClient *mongo.Client
}

type ToDoRepository interface {
	AddTask(task domain.Task) error
}

func NewToDoRepo(dbClient *mongo.Client) *toDoRepo {
	return &toDoRepo{
		dbClient: dbClient,
	}
}

func (repo *toDoRepo) AddTask(task domain.Task) error {
	return nil
}
