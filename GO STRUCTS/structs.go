package main

import (
	"fmt"
	"strconv"

	"example.com/structs/user"
)

func main() {
	//creating an instance of a struct and passing the values inside it
	name := takeUserInfo("Enter your name: ")
	ageStr := takeUserInfo("Enter your age: ")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}
	newAdmin := user.NewAdmin(false, "Ben", 50)
	newAdmin.PrintAdmin()
	newUser, err := user.New(name, age)
	if err != nil {
		fmt.Println(err)
		return
	}
	newUser.PrintUserInfo()

}

func takeUserInfo(text string) string {
	var value string
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}
