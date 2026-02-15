package cmdmanager

import (
	"errors"
	"fmt"
)

type CmdManager struct {
}

func New() *CmdManager {
	return &CmdManager{}
}

func (fileManager CmdManager) WriteJSON(data any) error {
	fmt.Println("OutPut:", data)

	return nil
}

func (fileManager CmdManager) Readlines() ([]string, error) {
	var prices = []string{}

	fmt.Println("Please input the prices")
	for {
		var price string
		fmt.Print("Price")
		_, err := fmt.Scan(&price)
		if err != nil {
			return nil, errors.New("error while scan")
		}
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}
