package main

import "fmt"

func main() {
	var age int = 32
	adultYears := getAdultYears(&age)
	fmt.Println("Age:", age)
	fmt.Println("Adult years:", adultYears)

	mutateAge(&age)

	fmt.Println("Mutated Age", age)
}

//using pointers to pass by reference
func getAdultYears(age *int) int {
	return *age - 18
}

//using pointers to change value of a variable
func mutateAge(age *int) {
	*age = 20
}
