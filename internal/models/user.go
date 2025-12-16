package models

import "fmt"

type User struct {
	ID   int    `json:"userid"`
	Name string `json:"name"`
}

func (user *User) Greet() string {
	return fmt.Sprintf("Hello, I am %s", user.Name)
}
