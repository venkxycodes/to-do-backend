package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"to-do/config"
	"to-do/domain"
)

type toDoRepo struct {
	collection *mongo.Collection
}

type ToDoRepository interface {
	AddTask(task domain.Task) error
}

func NewToDoRepo(db *mongo.Client) *toDoRepo {
	return &toDoRepo{
		collection: db.Database(config.GetConfig().DbConfig.DBName).Collection("tasks"),
	}
}

func (repo *toDoRepo) AddTask(task domain.Task) error {
	_, err := repo.collection.InsertOne(context.Background(), task)
	return err
}
