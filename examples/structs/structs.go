package structs

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

func (u *user) clearUserNames() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthday string) (*user, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("First name, last name, and birthday are required.")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil
}

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

	appUser, err := newUser(firstName, lastName, birthday)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()
	appUser.clearUserNames()
	appUser.outputUserDetails()
}
