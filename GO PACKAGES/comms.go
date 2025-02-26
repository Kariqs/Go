package main

import (
	"example.com/car/shared"
)

func presentOptions() {
	shared.PrintText("Select a Car to Purchase")
	shared.PrintText("1.Toyota")
	shared.PrintText("2.Mercedes Benz")
	shared.PrintText("3.BMW")
	shared.PrintText("4.Exit")
}
