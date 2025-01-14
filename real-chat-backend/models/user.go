package models

type User struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}
