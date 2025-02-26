package main

import (
	"fmt"
	"math"
)

func main() {
	const infaltionRate float64 = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the number of years:  ")
	fmt.Scan(&years)

	FV, RFV := calculateFutureValue(investmentAmount, expectedReturnRate, years, infaltionRate)
	fmt.Printf("Future Value: %.2f\nFuture Value After Inflation: %.2f", FV, RFV)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years, infaltionRate float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow((1+(expectedReturnRate/100)), years)
	realFutureValue := futureValue / math.Pow((1+(infaltionRate/100)), years)
	return futureValue, realFutureValue
}
