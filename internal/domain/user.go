package domain

type User struct {
	User_id string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

type UserRepository interface {
	Create(user *User) error
}
