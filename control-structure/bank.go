package main

import (
	"fmt"
"github.com/Pallinder/go-randomdata"
	"example.com/bank/fileops"
)

var balanceFile string = "balance.txt"

func main() {
	fmt.Println("Welcome to GO Bank!!")
	fmt.Println("Reach us at 24/7:",randomdata.PhoneNumber())
	// var (
	// 	accountBalance float64
	// 	error          error
	// )
	accountBalance, error := fileops.GetFloatValueFromFile(balanceFile)

	if error != nil {
		fmt.Println("There was a error reading file or parsing so working with default value")
	}

	var option int = 0
	for option != 4 {

		printOptions()
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("Account balance is ", accountBalance, "\n")
			fileops.WriteFloatValuetoFile(accountBalance, balanceFile)
		case 2:
			fmt.Print("Input the amount to deposit:")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit > 0 {
				accountBalance += deposit
			} else {
				fmt.Println("deposit value should be more than 0")
			}
			fileops.WriteFloatValuetoFile(accountBalance, balanceFile)
		case 3:
			fmt.Print("Input the amount to withdraw:")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw > accountBalance {
				fmt.Println("\n", "Your account does not have sufficient balance")
			} else if withdraw < 0 {
				fmt.Println("withdraw value should be more than 0")
			} else if withdraw < accountBalance {
				accountBalance -= withdraw
			}
			fileops.WriteFloatValuetoFile(accountBalance, balanceFile)
		case 4:
			fmt.Println("Thank you for using our services!!")
		default:
			fmt.Println("Input correct value")
		}
	}
}
