package contract

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"to-do/domain"
)

type CreateTask struct {
	UserId    int64           `json:"user_id" bson:"user_id"`
	Name      string          `json:"name" binding:"required" bson:"name"`
	Deadline  int64           `json:"deadline" binding:"required" bson:"deadline"`
	Priority  domain.Priority `json:"priority" binding:"required,enum" bson:"priority"`
	Notes     string          `json:"notes" bson:"notes"`
	CreatedBy string          `json:"created_by" bson:"created_by" binding:"required"`
}

type UpdateTask struct {
	Id        primitive.ObjectID `json:"id" bson:"id" binding:"required"`
	UserId    int64              `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" binding:"required" bson:"name"`
	Notes     string             `json:"notes" binding:"required" bson:"notes"`
	Deadline  int64              `json:"deadline" binding:"required" bson:"deadline"`
	Priority  domain.Priority    `json:"priority" binding:"required,enum" bson:"priority"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by" binding:"required"`
}

func (c *CreateTask) Validate() map[string]string {
	errors := make(map[string]string)
	if len(c.Name) == 0 {
		errors["name"] = "err-task-name-could-not-be-empty"
	}
	if c.UserId == 0 {
		errors["user_id"] = "err-task-user-id-cannot-be-zero"
	}
	// Allow deadlines upto 7 days before from today
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
