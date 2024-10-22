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
	Id        primitive.ObjectID `json:"id" bson:"id" binding:"required"`
	UserId    int64              `json:"user_id" bson:"user_id" binding:"required"`
	Name      string             `json:"name" binding:"required" bson:"name"`
	Notes     string             `json:"notes" binding:"required" bson:"notes"`
	Deadline  int64              `json:"deadline" binding:"required,CheckValidDeadline" bson:"deadline"`
	Priority  domain.Priority    `json:"priority" binding:"required" bson:"priority"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by" binding:"required"`
}

var CheckValidDeadline validator.Func = func(fl validator.FieldLevel) bool {
	deadlineInt, ok := fl.Field().Interface().(int64)
	if !ok || deadlineInt < time.Now().AddDate(0, 0, -7).UnixMilli() {
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
}

func (c *CreateTask) Validate() map[string]string {
	errors := make(map[string]string)
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
