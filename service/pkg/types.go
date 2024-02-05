package pkg

import "todo/entity"

type TaskResponse struct {
	Task entity.Task
}

type ListRequest struct {
	UserID int
}
type ListResponse struct {
	Tasks []entity.Task
}

type CreateTaskRequest struct {
	Title               string
	Category            entity.Category
	AuthenticatedUserID int
}
