package functions

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3}
	numbersDoubled := transformNumbers(&numbers, double)
	fmt.Println("doubled:", numbersDoubled)
	numbersTripled := transformNumbers(&numbers, triple)
	fmt.Println("tripled:", numbersTripled)


}



// func transformNumbers(numbers *[]int, transformer func(int) int) []int {
// since functions accepting and returning complex type would be hard to put , those can be alised like type
func transformNumbers(numbers *[]int, transformer transformFn) []int {
	var transformedNumber []int = []int{}
	for _, value := range *numbers {
		transformedNumber = append(transformedNumber, transformer(value))
	}
	return transformedNumber
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
