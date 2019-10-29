func minDistance(word1, word2 string) int {
	m := len(word1)
	n := len(word2)

	// dp[i][j] 表示 word1[:i] 和 word2[:j] 的编辑距离
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 删 [插入a、跟删除b没差别]
			dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])

			replace := 1
			if word1[i-1] == word2[j-1] {
				replace = 0
			}

			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+replace)
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

