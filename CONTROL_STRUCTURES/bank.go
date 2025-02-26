package main

import "fmt"

func main() {

	var accountBalance float64 = 20000
	printDetails("Welcome to Go Bank")
	for {
		printDetails("1. Check Balance")
		printDetails("2. Deposit Money")
		printDetails("3. Withdraw Money")
		printDetails("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is:", accountBalance)
		} else if choice == 2 {
			depositAmount := getAmount("dep")
			money("dep", depositAmount, accountBalance)
		} else if choice == 3 {
			withdrawAmount := getAmount("wit")
			money("wit", withdrawAmount, accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	printDetails("Thanks for choosing Go Bank")
}

func printDetails(text string) {
	fmt.Println(text)
}

func getAmount(phrase string) float64 {
	var amount float64
	if phrase == "dep" {
		fmt.Print("Enter amount to deposit: ")
	} else if phrase == "wit" {
		fmt.Print("Enter amount to withdraw: ")
	}
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Are you a retard? How can you deposit a value less than 0?")
		return 0
	}
	return amount
}

func money(phrase string, amount, balance float64) {
	if phrase == "dep" {
		accountBalance := balance + amount
		fmt.Printf("You have deposited %.2f to your account.\nYour new balance is %.2f", amount, accountBalance)
	} else if phrase == "wit" {
		accountBalance := balance - amount
		fmt.Printf("You have withdrawn %.2f to your account.\nYour new balance is %.2f", amount, accountBalance)
	}
}
