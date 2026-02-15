package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	appUser, err := user.New(firstName, lastName, birthdate)

	// if you want to have struct variable instead of pointer
	// var user user = *appUser
	// fmt.Println("user struct variable:", user)

	if err != nil {
		//fmt.Println("Error:", err)
		panic(err)
	}

	appUser.OutoutUserDate("appUser data is:")
	appUser.ClearFirstLastName()
	appUser.OutoutUserDate("appUser data is:")

	//fmt.Println(firstName, lastName, birthdate)

	adminUser := user.NewAdmin("example@org.com", "123")
	adminUser.OutoutUserDate("adminUser data is:")
	adminUser.ClearFirstLastName()
	adminUser.OutoutUserDate("adminUser data is:")
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
