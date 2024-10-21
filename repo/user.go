package repo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"to-do/config"
	"to-do/domain"
)

type userRepo struct {
	collection *mongo.Collection
}

type UserRepository interface {
	AddNewUser(ctx *gin.Context, user *domain.User) error
	GetUserByUsername(ctx *gin.Context, username string) (*domain.User, error)
	GetAllUsers(ctx *gin.Context) ([]domain.User, error)
}

func NewUserRepository(db *mongo.Client) *userRepo {
	return &userRepo{
		collection: db.Database(config.GetConfig().DbConfig.DBName).Collection("users"),
	}
}

func (r *userRepo) AddNewUser(ctx *gin.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepo) GetUserByUsername(ctx *gin.Context, username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	var result domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	return &result, err
}

func (r *userRepo) GetAllUsers(ctx *gin.Context) ([]domain.User, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(map[string]interface{}{"user_id": -1})
	cursor, err := r.collection.Find(ctx, filter, opts)
	var users []domain.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, err
}
