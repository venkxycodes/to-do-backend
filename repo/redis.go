package repo

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"to-do/domain"
)

type userCache struct {
	rdb *redis.Client
}

type UserCache interface {
	SetUser(*domain.User) error
	GetUserByUserName(ctx gin.Context, username string) error
}

func NewUserCache(rdb *redis.Client) UserCache {
	return &userCache{
		rdb: rdb,
	}
}

func (c *userCache) SetUser(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (c *userCache) GetUserByUserName(ctx gin.Context, username string) error {
	//TODO implement me
	panic("implement me")
}
