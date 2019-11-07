func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)

	// dp[i][j] 表示 s[:i] 与 t[:j] 的子序列方案
	dp := make([][]int, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]int, tLen+1)
	}
	for i := 0; i < sLen; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < tLen; j++ {
		for i := j; i < sLen; i++ {
			if t[j] == s[i] {
				// s[i] 不删，等价与 s[:i-1] 与 t[:j-1] 的配对
				// s[i] 删掉，等价于 s[:i-1] 与 t[:j] 的配对
				dp[i+1][j+1] = dp[i][j] + dp[i][j+1]
			} else {
				// s[i] 只能删掉
				dp[i+1][j+1] = dp[i][j+1]
			}
		}
	}
	return dp[sLen][tLen]
}
