package models

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:password`
}

func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}
