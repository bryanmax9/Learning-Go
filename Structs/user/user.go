package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u User) OutputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin( email, password string) Admin {
	return Admin {
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "-----",
			createdAt: time.Now(),
		},
	}
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) {

	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("Firstname,Lastname, and Birthdate are required")
	}

	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}
