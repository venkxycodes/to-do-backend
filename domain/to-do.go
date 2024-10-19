package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Priority string
type State string

const (
	HIGH   Priority = "high"
	MEDIUM Priority = "medium"
	LOW    Priority = "low"

	Pending    State = "pending"
	InProgress State = "in_progress"
	Completed  State = "completed"
)

type Task struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Deadline     int64              `json:"deadline" bson:"deadline"`
	Priority     Priority           `json:"priority" bson:"priority"`
	Notes        string             `json:"notes" bson:"notes"`
	CurrentState State              `json:"current_state" bson:"current_state"`
	UpsertMeta   UpsertMeta         `json:"upsert_meta" bson:"upsert_meta"`
}

type UpsertMeta struct {
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
	CreatedBy string `json:"created_by" bson:"created_by"`
	UpdatedBy string `json:"updated_by" bson:"updated_by"`
}
