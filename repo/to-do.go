package repo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"to-do/config"
	"to-do/domain"
)

type toDoRepo struct {
	collection *mongo.Collection
}

type ToDoRepository interface {
	AddTask(ctx *gin.Context, task domain.Task) error
	EditTask(ctx *gin.Context, task domain.Task) error
}

func NewToDoRepo(db *mongo.Client) *toDoRepo {
	return &toDoRepo{
		collection: db.Database(config.GetConfig().DbConfig.DBName).Collection("tasks"),
	}
}

func (repo *toDoRepo) AddTask(ctx *gin.Context, task domain.Task) error {
	_, err := repo.collection.InsertOne(ctx, task)
	return err
}

func (repo *toDoRepo) EditTask(ctx *gin.Context, task domain.Task) error {
	filter := bson.M{"_id": task.Id}
	update := bson.M{"$set": task}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}
