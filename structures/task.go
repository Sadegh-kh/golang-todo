package structures

type Task struct {
	title    string
	category Category
	isDone   bool
	date     string
	UserId   int
}

var taskStorage []Task

func (t *Task) CreateTask(title, date string, category Category, isDone bool, UserId int) {
	t.title = title
	t.date = date
	t.category = category
	t.isDone = isDone
	t.UserId = UserId
	taskStorage = append(taskStorage, *t)

}

func GetTaskList(UserId int) []Task {
	userTasks := []Task{}
	for _, value := range taskStorage {
		if UserId == value.UserId {
			userTasks = append(userTasks, value)
		}
	}
	return userTasks
}
