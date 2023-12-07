package bank

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILE = "examples/bank/balance.txt"
const DEFAULT_VALUE = 0

func readBalanceFromFile() (float64, error) {
	balance, err := os.ReadFile(ACCOUNT_BALANCE_FILE)
	if err != nil {
		return DEFAULT_VALUE, errors.New("Failed to read balance file.")
	}

	balanceText := string(balance)
	balanceAmount, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return DEFAULT_VALUE, errors.New("Failed to parse stored balance value.")
	}

	return balanceAmount, nil
}

func writeBalanceToFile(balance float64) {
	formattedBalance := fmt.Sprintf("%.2f", balance)
	os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(formattedBalance), 0644)
}

func Run() {
	accountBalance, err := readBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("----------")
		// panic(err)
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("How much do you want to deposit? ")
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("You can't deposit a negative amount!")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("How much do you want to withdraw? ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount < 0 {
				fmt.Println("You can't withdraw a negative amount!")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("You don't have enough balance!")
			} else {
				accountBalance -= withdrawalAmount
				writeBalanceToFile(accountBalance)
				fmt.Printf("Your balance is %.2f\n", accountBalance)
			}
		case 4:
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
