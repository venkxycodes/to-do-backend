package repo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"to-do/config"
	"to-do/domain"
)

type taskRepository struct {
	collection *mongo.Collection
}

type TaskRepository interface {
	AddTask(ctx *gin.Context, task *domain.Task) error
	EditTask(ctx *gin.Context, task *domain.Task) error
	GetAllTasksForUser(ctx *gin.Context, userId int64) ([]domain.Task, error)
	GetTaskById(ctx *gin.Context, id primitive.ObjectID) (*domain.Task, error)
}

func NewToDoRepo(db *mongo.Client) TaskRepository {
	return &taskRepository{
		collection: db.Database(config.GetConfig().DbConfig.DBName).Collection("tasks"),
	}
}

func (repo *taskRepository) AddTask(ctx *gin.Context, task *domain.Task) error {
	_, err := repo.collection.InsertOne(ctx, task)
	return err
}

func (repo *taskRepository) EditTask(ctx *gin.Context, task *domain.Task) error {
	filter := bson.M{"_id": task.Id}
	update := bson.M{"$set": task}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *taskRepository) GetAllTasksForUser(ctx *gin.Context, userId int64) ([]domain.Task, error) {
	filter := bson.M{"user_id": userId}                                         // Filter by user_id
	findOptions := options.Find().SetSort(bson.D{{Key: "deadline", Value: -1}}) // Sort by deadline in descending order
	cursor, err := repo.collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var tasks []domain.Task
	err = cursor.All(ctx, &tasks)
	return tasks, nil
}

func (repo *taskRepository) GetTaskById(ctx *gin.Context, id primitive.ObjectID) (*domain.Task, error) {
	var task *domain.Task
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	return task, err
}
