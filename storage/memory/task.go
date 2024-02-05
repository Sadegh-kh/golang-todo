package memory

import (
	"todo/entity"
)

type TaskStore struct {
	Tasks []entity.Task
}

func NewTaskStorege() *TaskStore {
	return &TaskStore{
		Tasks: []entity.Task{},
	}
}

func (t *TaskStore) StorageLen() int {
	return len(t.Tasks)
}
func (t *TaskStore) SaveTask(task entity.Task) (entity.Task, error) {
	t.Tasks = append(t.Tasks, task)
	return task, nil
}
func (t *TaskStore) GetListTask(UserID int) ([]entity.Task, error) {
	var userTasks []entity.Task
	for _, value := range t.Tasks {
		if UserID == value.UserID {
			userTasks = append(userTasks, value)
		}
	}
	return userTasks, nil
}
