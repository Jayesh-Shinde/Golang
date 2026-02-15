package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	transformerBy4 := getTransformerByFactor(4)
	transformerBy5 := getTransformerByFactor(5)

	numbers4 := transformNumbers(&numbers, transformerBy4)
	numbers5 := transformNumbers(&numbers, transformerBy5)

	fmt.Println("numbers4", numbers4)
	fmt.Println("numbers5", numbers5)

	result := factorial(5)
	fmt.Println("factorial:", result)
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}

func getTransformerByFactor(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
