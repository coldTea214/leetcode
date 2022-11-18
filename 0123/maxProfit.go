package main

import "fmt"

// 后置题 0188
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	//        3  2  6  5  0 3
	// buy1  -3 -2 -2 -2  0 0
	// sell1  0  0  4  4  4 4
	// buy2  -3 -2 -2 -1  4 4
	// sell2  0  0  4  4  4 7
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 7
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
}
