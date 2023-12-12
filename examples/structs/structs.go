package structs

import (
	"fmt"

	"github.com/mheob/udemy-go/examples/structs/user"
)

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}

func Run() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthday := getUserData("Please enter your birthday (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthday)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserNames()
	appUser.OutputUserDetails()
}
