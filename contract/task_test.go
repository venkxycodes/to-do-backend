package contract

import (
	"reflect"
	"testing"
	"time"
	"to-do/domain"
)

func TestCreateTask_Validate(t *testing.T) {
	currentTime := time.Now().UnixMilli()
	tests := []struct {
		name string
		c    *CreateTask
		want map[string]string
	}{
		{
			name: "test valid request",
			c: &CreateTask{
				UserName:  "venkxy30",
				Name:      "new task",
				Deadline:  currentTime,
				Priority:  domain.LOW,
				Notes:     "notes",
				CreatedBy: "venkat",
			},
			want: map[string]string{},
		},
		{
			name: "test invalid request",
			c: &CreateTask{
				UserName:  "",
				Name:      "",
				Deadline:  1629591318000,
				Priority:  "new priority",
				CreatedBy: "",
			},
			want: map[string]string{
				"user_name":  "err-user-name-cannot-be-empty",
				"name":       "err-task-name-could-not-be-empty",
				"deadline":   "err-task-deadline-cannot-be-before-7-days-from-today",
				"priority":   "err-task-priority-invalid",
				"created_by": "err-task-created-by-could-not-be-empty",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
