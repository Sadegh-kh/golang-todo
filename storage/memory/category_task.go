package memory

type CategoryTask struct {
	*Category
	*Task
}

func NewTaskCategoryStorege(taskStorage *Task, categoryStorage *Category) *CategoryTask {
	return &CategoryTask{
		Task:     taskStorage,
		Category: categoryStorage,
	}
}