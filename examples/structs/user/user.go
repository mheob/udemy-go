package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

type Admin struct {
	// User     User
	User
	email    string
	password string
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

func (u *User) ClearUserNames() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("First name, last name, and birthday are required.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email and password are required.")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthday:  "00/00/0000",
		},
	}, nil
}
