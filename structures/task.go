package structures

type Task struct {
	title    string
	category Category
	isDone   bool
	date     string
}

func (t *Task) CreateTask(title, date string, category Category, isDone bool) {
	t.title = title
	t.date = date
	t.category = category
	t.isDone = isDone
}
