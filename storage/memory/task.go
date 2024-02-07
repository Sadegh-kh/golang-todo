package memory

import (
	"todo/entity"
)

type Task struct {
	Tasks []entity.Task
}

func NewTaskStorege() *Task {
	return &Task{
		Tasks: []entity.Task{},
	}
}

func (t *Task) StorageLen() int {
	return len(t.Tasks)
}
func (t *Task) SaveTask(task entity.Task) (entity.Task, error) {
	t.Tasks = append(t.Tasks, task)
	return task, nil
}
func (t *Task) GetListTask(UserID int) ([]entity.Task, error) {
	var userTasks []entity.Task
	for _, value := range t.Tasks {
		if UserID == value.UserID {
			userTasks = append(userTasks, value)
		}
	}
	return userTasks, nil
}
