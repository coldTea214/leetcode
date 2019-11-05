func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	// dp[i] 表示 s[:i] 的编码个数
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			dp[i+1] = dp[i]
		} else if s[i-1] != '1' && s[i-1] != '2' {
			return 0
		}

		if s[i-1] == '1' ||
			s[i-1] == '2' && '0' <= s[i] && s[i] <= '6' {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[len(s)]
}
