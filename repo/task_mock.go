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

func (t *TaskRepoMock) AddTask(ctx *gin.Context, task *domain.Task) error {
	args := t.Called(ctx, task)
	return args.Error(0)
}

func (t *TaskRepoMock) EditTask(ctx *gin.Context, task *domain.Task) error {
	args := t.Called(ctx, task)
	return args.Error(0)
}

func (t *TaskRepoMock) GetAllTasksForUser(ctx *gin.Context, userId int64) ([]domain.Task, error) {
	args := t.Called(ctx, userId)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (t TaskRepoMock) GetTaskById(ctx *gin.Context, id primitive.ObjectID) (*domain.Task, error) {
	args := t.Called(ctx, id)
	return args.Get(0).(*domain.Task), args.Error(1)
}
