package main

import (
	"fmt"

	"example.com/car/desc"
	"example.com/car/shared"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	shared.PrintText("Welcome to our show room")
	randomData := randomdata.PhoneNumber()
	shared.PrintText(randomData)
	for {
		presentOptions()
		fmt.Print("Your Choice: ")
		var value int
		fmt.Scan(&value)
		desc.GetCarDescription(value)
	}

}
