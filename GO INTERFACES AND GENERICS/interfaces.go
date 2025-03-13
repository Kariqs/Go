package main

import (
	"example.com/interfaces/admin"
	"example.com/interfaces/user"
)

type printer interface {
	Print()
}

func main() {
	newUser := user.New("Bena", 25)
	printUser(newUser)

	newAdmin := admin.New(true, "Admin", 30)
	printUser(newAdmin)
}

func printUser(data printer) {
	data.Print()
}
