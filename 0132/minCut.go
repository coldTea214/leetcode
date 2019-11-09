func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// dp[i] = minCut(s[:i])
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = i - 1
	}

	for i := 0; i < n+1; i++ {
		// 以 i 为中心，j 为长度向左右扩展
		for j := 0; 0 <= i-j && i+j < n && s[i-j] == s[i+j]; j++ {
			if dp[i-j]+1 < dp[i+j+1] {
				dp[i+j+1] = dp[i-j] + 1
			}
		}

		// 以 i 和 i+1 中间的空为中心，向左右扩展
		for j := 1; 0 <= i-j+1 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			if dp[i-j+1]+1 < dp[i+j+1] {
				dp[i+j+1] = dp[i-j+1] + 1
			}
		}
	}

	return dp[n]
}
