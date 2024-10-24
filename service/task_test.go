package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
	"time"
	"to-do/contract"
	"to-do/domain"
	"to-do/repo"
)

func Test_toDoService_CreateTask(t *testing.T) {
	type fields struct {
		toDoRepo    repo.TaskRepoMock
		userService UserServiceMock
	}
	type args struct {
		ctx  *gin.Context
		task *contract.CreateTask
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test non identified user",
			fields: fields{
				toDoRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx:  ctx,
				task: &contract.CreateTask{UserName: "user"},
			},
			wantErr: true,
		},
		{
			name: "test identified user, add to repo err",
			fields: fields{
				toDoRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx:  ctx,
				task: &contract.CreateTask{UserName: "user"},
			},
			wantErr: true,
		},
		{
			name: "test identified user, add to repo success",
			fields: fields{
				toDoRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx:  ctx,
				task: &contract.CreateTask{UserName: "user"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &taskService{
				taskRepo:    &tt.fields.toDoRepo,
				userService: &tt.fields.userService,
			}
			if tt.name == "test non identified user" {
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, nil).Once()
			}
			if tt.name == "test identified user, add to repo err" {
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, fmt.Errorf("err-user-exists")).Once()
				tt.fields.toDoRepo.On("AddTask", tt.args.ctx, mock.Anything).Return(fmt.Errorf("err")).Once()
			}
			if tt.name == "test identified user, add to repo success" {
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, fmt.Errorf("err-user-exists")).Once()
				tt.fields.toDoRepo.On("AddTask", tt.args.ctx, mock.Anything).Return(nil).Once()
			}
			err := ts.CreateTask(tt.args.ctx, tt.args.task)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_taskService_UpdateTask(t *testing.T) {
	type fields struct {
		taskRepo    repo.TaskRepoMock
		userService UserServiceMock
	}
	type args struct {
		ctx  *gin.Context
		task *contract.UpdateTask
	}
	timestamp := time.Now().UnixMilli()
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test non identified user",
			fields: fields{
				taskRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx:  ctx,
				task: &contract.UpdateTask{UserName: "user", Deadline: timestamp},
			},
			wantErr: true,
		},
		{
			name: "test identified user, get task by id err",
			fields: fields{
				taskRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx:  ctx,
				task: &contract.UpdateTask{Id: primitive.NewObjectID(), UserName: "user"},
			},
			wantErr: true,
		},
		{
			name: "test identified user, get task by id nil task",
			fields: fields{
				taskRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx:  ctx,
				task: &contract.UpdateTask{Id: primitive.NewObjectID(), UserName: "user"},
			},
			wantErr: true,
		},
		{
			name: "test identified user, get task by id success, user id vs repo task user id mismatch",
			fields: fields{
				taskRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx:  ctx,
				task: &contract.UpdateTask{Id: primitive.NewObjectID(), UserName: "user"},
			},
			wantErr: true,
		},
		{
			name: "test identified user, edit task err",
			fields: fields{
				taskRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx: ctx,
				task: &contract.UpdateTask{
					Id:        primitive.NewObjectID(),
					UserName:  "user",
					Name:      "name",
					Notes:     "notes",
					Deadline:  timestamp,
					Priority:  domain.HIGH,
					UpdatedBy: "venkat",
				},
			},
			wantErr: true,
		},
		{
			name: "test identified user, edit task success",
			fields: fields{
				taskRepo:    repo.TaskRepoMock{},
				userService: UserServiceMock{},
			},
			args: args{
				ctx: ctx,
				task: &contract.UpdateTask{
					Id:        primitive.NewObjectID(),
					UserName:  "user",
					Name:      "name",
					Notes:     "notes",
					Deadline:  timestamp,
					Priority:  domain.HIGH,
					UpdatedBy: "venkat",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &taskService{
				taskRepo:    &tt.fields.taskRepo,
				userService: &tt.fields.userService,
			}
			if tt.name == "test non identified user" {
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, nil).Once()
			}
			if tt.name == "test identified user, get task by id err" {
				var taskModel *domain.Task
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, fmt.Errorf("err-user-already-exists")).Once()
				tt.fields.taskRepo.On("GetTaskById", tt.args.ctx, tt.args.task.Id).Return(taskModel, fmt.Errorf("err")).Once()
			}
			if tt.name == "test identified user, get task by id nil task" {
				var taskModel *domain.Task
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, fmt.Errorf("err-user-already-exists")).Once()
				tt.fields.taskRepo.On("GetTaskById", tt.args.ctx, tt.args.task.Id).Return(taskModel, nil).Once()
			}
			if tt.name == "test identified user, get task by id success, user id vs repo task user id mismatch" {
				taskModel := &domain.Task{
					UserId: 2,
				}
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, fmt.Errorf("err-user-already-exists")).Once()
				tt.fields.taskRepo.On("GetTaskById", tt.args.ctx, tt.args.task.Id).Return(taskModel, nil).Once()
			}
			if tt.name == "test identified user, edit task err" {
				taskModel := &domain.Task{
					UserId: 3,
				}
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, fmt.Errorf("err-user-already-exists")).Once()
				tt.fields.taskRepo.On("GetTaskById", tt.args.ctx, tt.args.task.Id).Return(taskModel, nil).Once()
				tt.fields.taskRepo.On("EditTask", tt.args.ctx, mock.Anything).Return(fmt.Errorf("err")).Once()
			}
			if tt.name == "test identified user, edit task success" {
				taskModel := &domain.Task{
					UserId: 3,
				}
				tt.fields.userService.On("GetUserIdByUserName", tt.args.task.UserName).Return(3, fmt.Errorf("err-user-already-exists")).Once()
				tt.fields.taskRepo.On("GetTaskById", tt.args.ctx, tt.args.task.Id).Return(taskModel, nil).Once()
				tt.fields.taskRepo.On("EditTask", tt.args.ctx, mock.Anything).Return(nil).Once()
			}
			err := ts.UpdateTask(tt.args.ctx, tt.args.task)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
