package contract

import "to-do/domain"

type CreateTask struct {
	Name      string          `json:"name" binding:"required" bson:"name"`
	Deadline  int64           `json:"deadline" binding:"required" bson:"deadline"`
	Priority  domain.Priority `json:"priority" bson:"priority"`
	Notes     string          `json:"notes" bson:"notes"`
	CreatedBy string          `json:"created_by" bson:"created_by"`
}
