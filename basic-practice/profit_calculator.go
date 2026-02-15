package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, expense, taxRate, error := getInput()
	if error != nil {
		fmt.Println("You can not enter any value as zero or negative")
		return
	}
	profitBeforeTax, profitAfterTax, ratio := calculateProfit(revenue, expense, taxRate)
	writeResultToFile(profitBeforeTax, profitAfterTax, ratio)
	fmt.Printf("Profit before tax, after tax and ratio are %.0f,%0.1f,%0.2f \n", profitBeforeTax, profitAfterTax, ratio)
	formattedRatio := fmt.Sprintf("Ratio with formatted value %0.2f", ratio)
	fmt.Println("New way of formatted value", formattedRatio)
	//fmt.Println("Profit before tax, after tax and ratio are ", profitBeforeTax, profitAfterTax, ratio)
}

func writeResultToFile(profitBeforeTax, profitAfterTax, ratio float64) {
	data := fmt.Appendf(nil, "Profit before tax, after tax and ratio are %.0f,%0.1f,%0.2f \n",
		profitBeforeTax, profitAfterTax, ratio)
	os.WriteFile("Result.txt", data, 0644)
}

func getInput() (float64, float64, float64, error) {
	var revenue, expense, taxRate float64
	fmt.Print("Please input the revenue:")
	fmt.Scan(&revenue)
	fmt.Print("Please input the expense:")
	fmt.Scan(&expense)
	fmt.Print("Please input the tax rate in percentage:")
	fmt.Scan(&taxRate)
	if revenue <= 0 || expense <= 0 || taxRate <= 0 {
		return revenue, expense, taxRate, errors.New("One of the value is zero or negative")
	}
	return revenue, expense, taxRate, nil
}

func calculateProfit(revenue, expense, taxRate float64) (float64, float64, float64) {
	profitBeforeTax := revenue - expense
	profitAfterTax := profitBeforeTax * (1 - taxRate/100)
	ratio := profitBeforeTax / profitAfterTax
	return profitBeforeTax, profitAfterTax, ratio
}
