package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// Calculate the "Earnings before tax" (EBT), "Earnings after tax" (profit) and the ratio
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}
