package main

import "fmt"

type transformFn func(int) int

func main() {
	//function as a parameter
	numbers := []int{1, 2, 3, 4, 5}
	transformedNumbers := transformNumbers(&numbers, getTransformerFunction(5))

	fmt.Println("Transformed slice:", transformedNumbers)

	sum := sumup(1, 10, 20, 30, 40, 50)
	fmt.Println(sum)

	//sumup using a slice. slices -> variadic func
	sliceSum := sumup(1, numbers...)
	fmt.Println(sliceSum)

	factor := factorial(6)
	fmt.Println("Factorial of n:", factor)

}

//Returning functions as value
func getTransformerFunction(factor int) transformFn {

	//Returning an anonymous function instead of the double function
	return func(number int) int {
		return number * factor
	}

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	var doubledNumbers []int
	for _, val := range *numbers {
		doubledNumbers = append(doubledNumbers, transform(val))
	}
	return doubledNumbers
}

// func double(val int) int {
// 	return val * 2
// }

// func quadriple(val int) int {
// 	return val * 4
// }

//Recursion
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

//Variadic function
func sumup(startingSum int, numbers ...int) int {
	sum := startingSum

	for _, val := range numbers {
		sum += val
	}

	return sum
}

/*
In this place, we learnt how to do the following:
1. Use functions as parameters
2. Use function as return values
3. Create and use anonymous functions - Functions with no name.
4. The concept of recursion - A function calling itself
5. Closure - An inner function using a variable defined in a parent function.const
6. Variadic function - A function that takes any number of variables with a specified type
7. Using slices as parameters for variadic functions with the three dots notation
*/
