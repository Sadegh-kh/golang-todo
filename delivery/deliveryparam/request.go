package deliveryparam

import "todo/entity"

type Request struct {
	Command string
	Task    CreateTask
}
type CreateTask struct {
	Title    string
	Category entity.Category
}

