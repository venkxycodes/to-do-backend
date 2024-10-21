package repo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"to-do/config"
	"to-do/domain"
)

type toDoRepo struct {
	collection *mongo.Collection
}

type ToDoRepository interface {
	AddTask(ctx *gin.Context, task *domain.Task) error
	EditTask(ctx *gin.Context, task *domain.Task) error
	GetAllTasksForUser(ctx *gin.Context, userId int64) ([]domain.Task, error)
	GetTaskById(ctx *gin.Context, id primitive.ObjectID) (*domain.Task, error)
}

func NewToDoRepo(db *mongo.Client) *toDoRepo {
	return &toDoRepo{
		collection: db.Database(config.GetConfig().DbConfig.DBName).Collection("tasks"),
	}
}

func (repo *toDoRepo) AddTask(ctx *gin.Context, task *domain.Task) error {
	_, err := repo.collection.InsertOne(ctx, task)
	return err
}

func (repo *toDoRepo) EditTask(ctx *gin.Context, task *domain.Task) error {
	filter := bson.M{"_id": task.Id}
	update := bson.M{"$set": task}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *toDoRepo) GetAllTasksForUser(ctx *gin.Context, userId int64) ([]domain.Task, error) {
	findOptions := options.Find()
	findOptions.SetSort(map[string]interface{}{"user_id": userId, "deadline": -1})
	cursor, err := repo.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var tasks []domain.Task
	err = cursor.All(ctx, &tasks)
	return tasks, nil
}

func (repo *toDoRepo) GetTaskById(ctx *gin.Context, id string) (*domain.Task, error) {
	var task *domain.Task
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	return task, err
}
