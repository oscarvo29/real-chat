package models

import "github.com/google/uuid"

type User struct {
	Uuid     uuid.UUID `json:"uuid"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}
