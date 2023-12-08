package bank

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/mheob/udemy-go/utils"
)

const ACCOUNT_BALANCE_FILE = "examples/bank/balance.txt"

func Run() {
	accountBalance, err := utils.GetFloatFromFile(ACCOUNT_BALANCE_FILE)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("----------")
		// panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		printOptions()

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
			utils.WriteFloatToFile(ACCOUNT_BALANCE_FILE, accountBalance)
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
				utils.WriteFloatToFile(ACCOUNT_BALANCE_FILE, accountBalance)
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
