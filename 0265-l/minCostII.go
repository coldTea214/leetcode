package main

import "fmt"

// 前置题 256
func minCostII(costs [][]int) int {
	n := len(costs[0])
	dp := costs[0]
	for _, cost := range costs[1:] {
		dpNew := make([]int, n)
		for j, c := range cost {
			dpNew[j] = 1 << 31
			for k := 1; k < n; k++ {
				if dp[(j+k)%n] < dpNew[j] {
					dpNew[j] = dp[(j+k)%n]
				}
			}
			dpNew[j] += c
		}
		dp = dpNew
	}

	min := dp[0]
	for i := 1; i < n; i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}
	return min
}

func main() {
	fmt.Println(minCostII([][]int{{1, 5, 3}, {2, 9, 4}}))
}
