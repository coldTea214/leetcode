func isMatch(s, p string) bool {
	sLen := len(s)
	pLen := len(p)

	// dp[i][j] 代表了 s[:i] 能否与 p[:j] 匹配
	dp := make([][]bool, sLen+1)
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}

	dp[0][0] = true
	// 处理 a*b*c* 场景
	for j := 1; j < pLen && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sLen; i++ {
		for j := 0; j < pLen; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				// 默认不会出现 p = "*" 这种非正则
				if p[j-1] != s[i] && p[j-1] != '.' {
					// p[j] 与 s[i] 不匹配，则 "x*" 解释为 ""
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					// p[j] 与 s[i] 匹配
					dp[i+1][j+1] =
						dp[i+1][j-1] || // "x*" 解释为 "", x and x*x, i不动，j两个作废
							dp[i+1][j] || // "x*" 解释为 "x", xx and x*x, i不动，j *作废
							dp[i][j+1] // "x*" 解释为 "x...", xx and x*, j不动，i还匹配同一个*
				}
			}
		}
	}

	return dp[sLen][pLen]
}
