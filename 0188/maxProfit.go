package main

import "fmt"

// 前置题 0123
func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// buy[i]  表示持有第i次交易股票时的最大收益
	// sell[i] 表示不持有第i次交易股票时的最大收益
	// buy[1] = -3, 因为买入股票导致收益为负
	buy, sell := make([]int, k+1), make([]int, k+1)
	for i := 0; i <= k; i++ {
		buy[i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 7
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
