package main

import "fmt"

//Type aliases
type ratingMap map[string]float64
type namesSlice []string

func main() {
	//Creating a slice using the make() function in Go
	names := make(namesSlice, 2)
	names = append(names, "Benard", "Dancun")
	fmt.Println(names)

	//Creating a map using the make() function in Go
	courseRatings := make(ratingMap, 3)
	courseRatings["go"] = 4.5
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9
	fmt.Println(courseRatings)
}
