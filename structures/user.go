package structures

type User struct {
	ID       int
	name     string
	Email    string
	Password string
}

var userStorage []User

func (u *User) CreateUser(name, email, password string) {
	u.ID = len(userStorage) + 1
	u.name = name
	u.Email = email
	u.Password = password

}

func GetUser(email string) User {
	for _, value := range userStorage {
		if email == value.Email {
			return value
		}
	}
	return User{}
}

func CheckPass(email, password string) bool {
	for _, value := range userStorage {
		if value.Email == email && value.Password == password {
			return true
		}
	}
	return false
}

func UserExist(email string) bool {
	for _, value := range userStorage {
		// email is the primary key
		if email == value.Email {
			return true
		}
	}
	return false
}

func (newUser User) AppendToStorage() {
	userStorage = append(userStorage, newUser)
}
