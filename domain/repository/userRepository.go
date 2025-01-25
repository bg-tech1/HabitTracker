package repository

type User struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRepository interface {
	LoginUser(id string) (*User, error)
	RegisterUser(id string, password string) error
	GetUser(id string) (*User, error)
	CreateSession(userID string, sessionID string) error
	GetUserId(sessionID string) (string, error)
}
