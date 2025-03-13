package admin

import (
	"fmt"

	"example.com/interfaces/user"
)

type Admin struct {
	user    user.User
	isAdmin bool
}

func New(isAdmin bool, name string, age int) *Admin {
	return &Admin{
		user: user.User{
			Name: name,
			Age:  age,
		},
		isAdmin: isAdmin,
	}
}

func (user *Admin) Print() {
	fmt.Println("You are an Admin. Your username is", user.user.Name, "and you are", user.user.Age, "years old.")
}
