package repo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"to-do/config"
	"to-do/domain"
)

type userRepository struct {
	collection *mongo.Collection
	counters   *mongo.Collection
}

type UserRepository interface {
	AddNewUser(ctx *gin.Context, user *domain.User) error
	GetUserByUserId(ctx *gin.Context, userId int64) (*domain.User, error)
	GetUserByUsername(ctx *gin.Context, username string) (*domain.User, error)
	GetAllUsers(ctx *gin.Context) ([]domain.User, error)
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		collection: db.Database(config.GetConfig().DbConfig.DBName).Collection("users"),
		counters:   db.Database(config.GetConfig().DbConfig.DBName).Collection("counters"),
	}
}

func (r *userRepository) NextUserID(ctx *gin.Context) (int64, error) {
	var counter struct {
		Value int64 `bson:"value"`
	}
	result := r.counters.FindOneAndUpdate(ctx, bson.M{"_id": "users"}, bson.M{"$inc": bson.M{"value": int64(1)}}, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))
	if err := result.Decode(&counter); err != nil {
		return 0, err
	}
	return counter.Value, nil
}

func (r *userRepository) AddNewUser(ctx *gin.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) GetUserByUserId(ctx *gin.Context, userId int64) (*domain.User, error) {
	filter := bson.M{"user_id": userId}
	var result domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	return &result, err
}

func (r *userRepository) GetUserByUsername(ctx *gin.Context, username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	var result domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	return &result, err
}

func (r *userRepository) GetAllUsers(ctx *gin.Context) ([]domain.User, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(map[string]interface{}{"user_id": -1})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var users []domain.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, err
}
