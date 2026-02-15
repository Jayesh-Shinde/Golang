package main

import "fmt"

func main() {
	var age int = 34
	agePointer := &age

	fmt.Println("Age value from orignal variable is:", age)
	fmt.Println("Age pointer address", agePointer)
	fmt.Println("Age value from it's pointer")

	adultAge := calculateAdultAge(agePointer) // here the pointer is passed to the function
	// dereferencing happens inside the function
	// new copy age will not be created

	fmt.Println("Adult age is: ", adultAge)

	calculateAdultAgeInPlace(agePointer)

	fmt.Println("Orignal age variable is changed:", age)
}

func calculateAdultAge(age *int) int {
	return *age - 18
}

func calculateAdultAgeInPlace(age *int) {
	*age = *age - 18
}
