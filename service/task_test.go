package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
	"to-do/contract"
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
