// 递归
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return isInterleaveHelper(s1, s2, s3)
}

func isInterleaveHelper(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}

	if s1[0] == s3[0] {
		if isInterleaveHelper(s1[1:], s2, s3[1:]) {
			return true
		}
	}
	if s2[0] == s3[0] {
		return isInterleaveHelper(s1, s2[1:], s3[1:])
	}
	return false
}

// 迭代
func isInterleave2(s1 string, s2 string, s3 string) bool {
    m, n, t := len(s1), len(s2), len(s3)
    if (m + n) != t {
        return false
    }
    dp := make([][]bool, m + 1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]bool, n + 1)
    }
    dp[0][0] = true
    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            p := i + j - 1
            if i > 0 {
                dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
            }
            if j > 0 {
                dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
            }
        }
    }
    return dp[m][n]
}
