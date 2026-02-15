package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	buyPrice := prices[0]
	profit := 0
	for _, sellPrice := range prices[1:] {
		if sellPrice < buyPrice {
			buyPrice = sellPrice
		}
		profit = max(profit, sellPrice-buyPrice)
	}
	return profit
}
