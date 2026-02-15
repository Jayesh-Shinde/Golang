package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println("Sum:", sumUp(numbers))
	fmt.Println("Sum II:", sumUpII(1, 2, 3)) //1 here is value of initialValue, numbers are 2,3
	fmt.Println("Sum II:", sumUpII(1, 2, 3, 4))
	fmt.Println("Sum II:(split array to parameters) ", sumUpII(0, numbers...))
}

func sumUpII(initialValue int, numbers ...int) int {
	var sum = initialValue
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func sumUp(numbers []int) int {
	var sum = 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
