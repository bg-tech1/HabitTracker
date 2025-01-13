package repository

type User struct {
	Id       string
	Password string
	Name     string
}

type UserRepository interface {
	LoginUser(id string, password string) (bool, error)
	RegisterUser(id string, password string) error
	GetUser(id string) (*User, error)
}
