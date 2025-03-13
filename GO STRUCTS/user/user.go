package user

import (
	"errors"
	"fmt"
)

// creating a struct
type User struct {
	name string
	age  int
}

// struct embedding
type Admin struct {
	isAdmin bool
	user    User
}

// outputting a struct
func (user *User) PrintUserInfo() {
	fmt.Println("User Details")
	fmt.Printf("Username: %v\nAge: %v", user.name, user.age)
}

// mutating a struct using a method
func (user *User) ClearName() {
	user.name = ""
}

func NewAdmin(isAdmin bool, name string, age int) *Admin {
	return &Admin{
		isAdmin: isAdmin,
		user: User{
			name: name,
			age:  age},
	}
}

func (admin Admin) PrintAdmin() {
	if admin.isAdmin {
		fmt.Println("This is an administrator")
	} else {
		fmt.Println("You are not an administrator")
	}
}

// constructor
func New(name string, age int) (*User, error) {
	if name == "" || age == 0 {
		return nil, errors.New("noma sana")
	}
	return &User{
		name: name,
		age:  age,
	}, nil
}
