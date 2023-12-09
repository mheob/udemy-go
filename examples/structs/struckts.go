package structs

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}

func Run() {
	appUser := user{
		firstName: getUserData("Please enter your first name: "),
		lastName:  getUserData("Please enter your last name: "),
		birthday:  getUserData("Please enter your birthday (MM/DD/YYYY): "),
		createdAt: time.Now(),
	}

	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthday)
}
