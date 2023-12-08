package profit_calculator

import (
	"errors"
	"fmt"
	"os"
)

const RESULTS_FILE = "exercises/profit_calculator/results.txt"

// storeResults writes the calculated EBT, profit, and ratio to a file.
// It takes three float64 arguments: ebt, profit, and ratio.
// The results are formatted as a string and written to the file specified by RESULTS_FILE.
// The file is created with read and write permissions for the owner and read permissions for others.
func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(RESULTS_FILE, []byte(results), 0644)
}

// printResults prints the calculated EBT, profit, and ratio.
// It takes three float64 arguments: ebt, profit, and ratio.
func printResults(ebt, profit, ratio float64) {
	fmt.Printf("Earnings before tax: %.1f\n", ebt)
	fmt.Printf("Earnings after tax: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

// getUserInput prompts the user with a question and returns the user's input as a float64.
func getUserInput(question string) (float64, error) {
	var userInput float64
	fmt.Print(question)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
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
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	storeResults(ebt, profit, ratio)
	printResults(ebt, profit, ratio)
}
