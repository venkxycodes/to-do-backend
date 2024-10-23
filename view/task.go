package view

import "to-do/domain"

type GetTasksResponse struct {
	Tasks []domain.Task `json:"tasks"`
}
