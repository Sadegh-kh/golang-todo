package storage

import "todo/structures"

type Storage interface {
	Create()
	Load()
	Save(u structures.User)
}
