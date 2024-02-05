package service

import (
	"fmt"
	"todo/entity"
	"todo/service/pkg"
	"todo/storage/memory"
)

type TaskServiceStorag interface {
	StorageLen() int
	SaveTask(t entity.Task) (entity.Task, error)
	GetListTask(userID int) (pkg.ListResponse, error)
}

type Task struct {
	storage TaskServiceStorag
}

func NewTaskService() Task {
	taskStorage := memory.NewTaskStorege()
	return Task{
		storage: &taskStorage,
	}
}

func (t Task) CreateTask(req pkg.CreateTaskRequest) (pkg.TaskResponse, error) {

	newTask := entity.Task{
		ID:       t.storage.StorageLen() + 1,
		Title:    req.Title,
		Category: req.Category,
		UserID:   req.AuthenticatedUserID,
		IsDone:   false,
	}
	newTask, err := t.storage.SaveTask(newTask)
	if err != nil {
		return pkg.TaskResponse{}, fmt.Errorf("error happen when save task to storage because: %v", err)
	}
	responseTask := pkg.TaskResponse{
		Task: newTask,
	}
	return responseTask, nil

}

func (t Task) GetListTask(req pkg.ListRequest) (pkg.ListResponse, error) {

	tasks, err := t.storage.GetListTask(req.UserID)
	if err != nil {
		return pkg.ListResponse{}, fmt.Errorf("can't give list tasks becuse: %v", err)
	}
	return tasks, nil

}
