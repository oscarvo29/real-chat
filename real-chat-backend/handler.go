package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var ActiveUsers []User

func (app *Config) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index page"))
}

func (app *Config) GetActiveUsers(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(ActiveUsers)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		panic(err)
	}
}

func (app *Config) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login have been hit")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usr User
	err = json.Unmarshal(body, &usr)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range ActiveUsers {
		if user.Name == usr.Name && user.Password == usr.Password {
			out, err := json.Marshal(usr)
			if err != nil {
				log.Fatal(err)
			}

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write(out)
			if err != nil {
				panic(err)
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("user-not-found"))
	if err != nil {
		panic(err)
	}

}

func (app *Config) SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup have been hit")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usr User
	err = json.Unmarshal(body, &usr)
	if err != nil {
		log.Fatal(err)
	}

	ActiveUsers = append(ActiveUsers, usr)
	out, err := json.Marshal(usr)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		panic(err)
	}
}
