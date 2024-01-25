package structures

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var UserStorage []User
var err error

func (u *User) CreateUser(Name, email, password string) {
	u.ID = len(UserStorage) + 1
	u.Name = Name
	u.Email = email
	u.Password = password

}

func GetUser(email string) User {
	for _, value := range UserStorage {
		if email == value.Email {
			return value
		}
	}
	return User{}
}

func CheckPass(email, password string) bool {
	for _, value := range UserStorage {
		if value.Email == email && value.Password == password {
			return true
		}
	}
	return false
}

func UserExist(email string) bool {
	for _, value := range UserStorage {
		// email is the primary key
		if email == value.Email {
			return true
		}
	}
	return false
}

func (newUser User) AppendToStorage() {
	UserStorage = append(UserStorage, newUser)
}
