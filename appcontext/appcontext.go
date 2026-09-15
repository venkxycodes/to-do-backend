package appcontext

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"to-do/config"
)

type appContext struct {
	mongoDbClient *mongo.Client
	redisClient   *redis.Client
}

var appCtx *appContext

func Init() error {
	client, err := newDbClient()
	if err != nil {
		return err
	}
	if err := ensureIndexes(client); err != nil {
		_ = client.Disconnect(context.Background())
		return err
	}
	appCtx = &appContext{mongoDbClient: client}
	return nil
}

func newDbClient() (*mongo.Client, error) {
	dbConfig := config.GetConfig().DbConfig
	appName := config.GetConfig().AppName
	if len(appName) == 0 {
		appName = "default"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()

	connectionString := dbConfig.GetConnectionString()

	client, err := mongo.
		Connect(ctx, options.Client().
			ApplyURI(connectionString).
			SetMaxPoolSize(30).
			SetMinPoolSize(10).
			SetMaxConnecting(10).
			SetRetryWrites(false))
	if err != nil {
		return nil, fmt.Errorf("create mongo client: %w", err)
	}

	// Ping the primary
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("ping mongo: %w", err)
	}

	return client, nil
}

func ensureIndexes(client *mongo.Client) error {
	db := client.Database(config.GetConfig().DbConfig.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Collection("users").Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "username", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "user_id", Value: 1}}, Options: options.Index().SetUnique(true)},
	})
	if err != nil {
		return fmt.Errorf("create user indexes: %w", err)
	}
	var highestUser struct {
		UserID int64 `bson:"user_id"`
	}
	err = db.Collection("users").FindOne(ctx, bson.D{}, options.FindOne().SetSort(bson.D{{Key: "user_id", Value: -1}})).Decode(&highestUser)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("read highest user id: %w", err)
	}
	_, err = db.Collection("counters").UpdateOne(ctx, bson.M{"_id": "users"}, bson.M{"$max": bson.M{"value": highestUser.UserID}}, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("initialize user id counter: %w", err)
	}
	_, err = db.Collection("tasks").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "deadline", Value: -1}},
	})
	if err != nil {
		return fmt.Errorf("create task indexes: %w", err)
	}
	return nil
}

func GetDBClient() *mongo.Client {
	return appCtx.mongoDbClient
}

func GetRedisClient() *redis.Client {
	return appCtx.redisClient
}
