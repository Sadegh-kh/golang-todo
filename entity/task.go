package entity

type Task struct {
	ID       int
	Title    string
	Category Category
	IsDone   bool
	UserID   int
}
