package main

import (
	"fmt"

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

	//Call the generic function
	result := add(1, 1)
	name := add("Bena ", "Kariuki")

	fmt.Println(result, name)
}

func printUser(data printer) {
	data.Print()
}

// Generics feature
func add[T int | float64 | string](a, b T) T {
	return a + b
}
