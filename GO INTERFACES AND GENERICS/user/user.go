package user

import "fmt"

type User struct {
	Name string
	Age  int
}

func New(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (user *User) Print() {
	fmt.Println("You are a user. Your username is", user.Name, "and you are", user.Age, "years old.")
}
