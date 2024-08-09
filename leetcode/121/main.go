package main

import "fmt"

func main() {
	nums := []int{8, 3, 6, 9, 4, 1, 4, 5, 2}
	result := maxProfit(nums)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit = max(profit, prices[i]-minPrice)
		}
	}

	return profit
}
