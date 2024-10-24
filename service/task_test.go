package service

import (
	"github.com/gin-gonic/gin"
	"testing"
	"to-do/contract"
	"to-do/repo"
)

func Test_toDoService_CreateTask(t1 *testing.T) {
	type fields struct {
		toDoRepo    repo.TaskRepository
		userService UserServiceMock
	}
	type args struct {
		ctx  *gin.Context
		task *contract.CreateTask
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &taskService{
				taskRepo:    tt.fields.toDoRepo,
				userService: tt.fields.userService,
			}
			if err := t.CreateTask(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t1.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
