package main

import (
	"fmt"
	"math"
)

func minCut(s string) int {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := range s {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true
		if i > 0 {
			isPalindrome[i][i-1] = true // 无物理含义, 方便下面dp迭代
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i+1][j-1]
		}
	}

	// dp[i] = minCut(s[:i+1])
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if isPalindrome[0][i] {
			continue
		}
		dp[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if isPalindrome[j+1][i] && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(minCut("aab"))
}
