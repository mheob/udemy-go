package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	// var investmentAmount, expectedReturnRate, years float64
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureInvestmentValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureInvestmentValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value is:", futureInvestmentValue)
	fmt.Println("Future real value is:", futureRealValue)
}
