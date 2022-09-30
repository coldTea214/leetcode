func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)

	// dp[i][j] == true 表示 s[:i] 和 p[:j] 匹配
	dp := make([][]bool, sLen+1)
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}

	dp[0][0] = true
	// 处理 *** 场景
	for j := 1; j <= pLen; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 0; i < sLen; i++ {
		for j := 0; j < pLen; j++ {
			if p[j] == s[i] || p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] =
					dp[i][j+1] || // * 匹配任意字符
						dp[i+1][j] // * 匹配空字符
			}
		}
	}

	return dp[sLen][pLen]
}
