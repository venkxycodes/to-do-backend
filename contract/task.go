package contract

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"to-do/domain"
)

type CreateTask struct {
	UserName  string          `json:"user_name" bson:"user_name" binding:"required"`
	Name      string          `json:"name" binding:"required" bson:"name"`
	Deadline  int64           `json:"deadline" binding:"required,CheckValidDeadline" bson:"deadline"`
	Priority  domain.Priority `json:"priority" binding:"required" bson:"priority"`
	Notes     string          `json:"notes" bson:"notes"`
	CreatedBy string          `json:"created_by" bson:"created_by" binding:"required"`
}

type UpdateTask struct {
	Id        primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	UserName  string             `json:"user_name" bson:"user_name" binding:"required"`
	Name      string             `json:"name" binding:"required" bson:"name"`
	Notes     string             `json:"notes" binding:"required" bson:"notes"`
	Deadline  int64              `json:"deadline" binding:"required,CheckValidDeadline" bson:"deadline"`
	Priority  domain.Priority    `json:"priority" binding:"required" bson:"priority"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by" binding:"required"`
}

type UpdateTaskStatus struct {
	TaskId    primitive.ObjectID `json:"task_id" bson:"task_id" binding:"required"`
	UserName  string             `json:"user_name" bson:"user_name" binding:"required"`
	State     domain.State       `json:"state" bson:"state" binding:"required"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by" binding:"required"`
}

var CheckValidDeadline validator.Func = func(fl validator.FieldLevel) bool {
	deadlineInt, ok := fl.Field().Interface().(int64)
	if !ok || deadlineInt < time.Now().AddDate(0, 0, -7).UnixMilli() {
		return false
	}
	return true
}

var CheckValidUserName validator.Func = func(fl validator.FieldLevel) bool {
	userName, ok := fl.Field().Interface().(string)
	if !ok || len(userName) < 8 {
		return false
	}
	return true
}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("CheckValidDeadline", CheckValidDeadline)
		if err != nil {
			return
		}
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("CheckValidUserName", CheckValidUserName)
		if err != nil {
			return
		}
	}
}

func (c *CreateTask) Validate() map[string]string {
	errors := make(map[string]string, 6)
	if len(c.Name) == 0 {
		errors["name"] = "err-task-name-could-not-be-empty"
	}
	if len(c.UserName) == 0 {
		errors["user_name"] = "err-user-name-cannot-be-empty"
	}
	// Prevent deadlines which are greater than 7 days from today
	if c.Deadline < time.Now().AddDate(0, 0, -7).UnixMilli() {
		errors["deadline"] = "err-task-deadline-cannot-be-before-7-days-from-today"
	}
	switch c.Priority {
	case domain.HIGH, domain.MEDIUM, domain.LOW:
	default:
		errors["priority"] = "err-task-priority-invalid"
	}
	if c.CreatedBy == "" {
		errors["created_by"] = "err-task-created-by-could-not-be-empty"
	}
	return errors
}

func (u *UpdateTask) Validate() map[string]string {
	errors := make(map[string]string)
	if _, err := primitive.ObjectIDFromHex(u.Id.Hex()); err != nil {
		errors["id"] = "err-task-id-invalid"
	}
	if len(u.Name) == 0 {
		errors["name"] = "err-task-name-could-not-be-empty"
	}
	if len(u.UserName) == 0 {
		errors["user_name"] = "err-user-name-cannot-be-empty"
	}
	// Prevent deadlines which are greater than 7 days from today
	if u.Deadline < time.Now().AddDate(0, 0, -7).UnixMilli() {
		errors["deadline"] = "err-task-deadline-cannot-be-before-7-days-from-today"
	}
	switch u.Priority {
	case domain.HIGH, domain.MEDIUM, domain.LOW:
	default:
		errors["priority"] = "err-task-priority-invalid"
	}
	if u.UpdatedBy == "" {
		errors["updated_by"] = "err-task-updated-by-could-not-be-empty"
	}
	return errors
}

func (ut *UpdateTaskStatus) Validate() map[string]string {
	errors := make(map[string]string)
	if _, err := primitive.ObjectIDFromHex(ut.TaskId.Hex()); err != nil {
		errors["id"] = "err-task-id-invalid"
	}
	if len(ut.UserName) == 0 {
		errors["user_name"] = "err-user-name-could-not-be-empty"
	}
	switch ut.State {
	case domain.InProgress, domain.Completed, domain.Pending:
	default:
		errors["state"] = "err-task-state-invalid"
	}
	if len(ut.UpdatedBy) == 0 {
		errors["updated_by"] = "err-task-updated-by-could-not-be-empty"
	}
	return errors
}
