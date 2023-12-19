package structures

type User struct {
	name     string
	Email    string
	Password string
}

func (u *User) CreateUser(name, email, password string) {
	u.name = name
	u.Email = email
	u.Password = password

}
