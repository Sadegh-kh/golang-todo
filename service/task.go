package service

import (
	"fmt"
	"todo/entity"
	"todo/storage/memory"
)

type TaskService interface {
	StorageLen() int
	SaveTask(t entity.Task) (entity.Task, error)
	GetListTask(userID int) ([]entity.Task, error)
}
type CategoryService interface {
	CheckCategoryByUserID(userID,categoryID int)bool
}

type Task struct {
	storage TaskService
	CategoryStorage CategoryService
}

func NewTaskService() Task {
	taskStorage := memory.NewTaskStorege()
	// categoryStorage:=memory.NewCategoryStorage()
	// categoryTaskStorage := memory.NewTaskCategoryStorege(taskStorage,categoryStorage)
	return Task{
		storage: taskStorage,
	}
}

func (t Task) CreateTask(req CreateTaskRequest) (TaskResponse, error) {

	newTask := entity.Task{
		ID:       t.storage.StorageLen() + 1,
		Title:    req.Title,
		Category: req.Category,
		UserID:   req.AuthenticatedUserID,
		IsDone:   false,
	}
	newTask, err := t.storage.SaveTask(newTask)
	if err != nil {
		return TaskResponse{}, fmt.Errorf("error happen when save task to storage because: %v", err)
	}
	responseTask := TaskResponse{
		Task: newTask,
	}
	return responseTask, nil

}

func (t Task) GetListTask(req ListRequest) (ListResponse, error) {

	tasks, err := t.storage.GetListTask(req.UserID)
	if err != nil {
		return ListResponse{}, fmt.Errorf("can't give list tasks becuse: %v", err)
	}
	return ListResponse{Tasks: tasks}, nil

}


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
