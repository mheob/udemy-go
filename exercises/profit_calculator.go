package exercises

import (
	"fmt"
)

// getUserInput prompts the user with a question and returns the user's input as a float64.
func getUserInput(question string) float64 {
	var userInput float64
	fmt.Print(question)
	fmt.Scan(&userInput)
	return userInput
}

// calculateFinancial calculates the "Earnings before tax" (EBT), "Earnings after tax" (profit), and the ratio.
// It takes the revenue, expenses, and tax rate as input parameters and returns the calculated EBT, profit, and ratio.
func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func Run() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	taxRate = getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.0f\n", ebt)
	fmt.Printf("Earnings after tax: %.0f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}
