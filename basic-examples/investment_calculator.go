package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.0
	var initialInvestment,years,annualInterestRate float64

	fmt.Print("Please input the initial investment amount:")
	fmt.Scan(&initialInvestment)

	fmt.Print("Please input the expected return rate:")
	fmt.Scan(&annualInterestRate)

	fmt.Print("Please input the number of years to invest:")
	fmt.Scan(&years)

	futureValue := initialInvestment * math.Pow((1+annualInterestRate/100), years)
	adjustedFutureValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future value is ", futureValue)
	fmt.Println("FutureAdjusted value is ", adjustedFutureValue)
}
