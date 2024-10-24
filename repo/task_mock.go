package repo

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"to-do/domain"
)

type TaskRepoMock struct {
	mock.Mock
}

func (t TaskRepoMock) AddTask(ctx *gin.Context, task *domain.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepoMock) EditTask(ctx *gin.Context, task *domain.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepoMock) GetAllTasksForUser(ctx *gin.Context, userId int64) ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepoMock) GetTaskById(ctx *gin.Context, id primitive.ObjectID) (*domain.Task, error) {
	//TODO implement me
	panic("implement me")
}
