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
	email    string
	password string
	User     // embedding User struct inside Admin struct
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			createdAt: time.Now(),
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "1/1/1",
		},
	}
}

func (u User) OutoutUserDate(initialText string) {
	fmt.Println(initialText,
		`firstName:`, u.firstName,
		`lastName: `, u.lastName,
		`birthday:`, u.birthdate,
		`createdAt:`, u.createdAt)
}

// below is not a constructor function but kind of a utility function to
// create a new use struct variable, benifit is you can add validation logic
// return pointer always as it allows to return nil in case of error
// instead of copy huge struct variable just return the memory address

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("one of the firstName,lastName or birthday should not have blank vale")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

//func (u *User) clearFirstLastName() { --> it is pointer receiver so value
// on original struct variable will be changed
//func (u User) clearFirstLastName() { --> it is value receiver so function
// receives a copy of the struct variable and orignal value will not be changed

func (u *User) ClearFirstLastName() {
	u.firstName = ""
	u.lastName = ""
}
