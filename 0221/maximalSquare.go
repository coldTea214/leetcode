func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	maxEdge := 0
	// dp[i][j] == 以 [i,j] 点为右下角点的最大正方形的边长
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxEdge = 1
		}
	}
	for j := 1; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxEdge = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
				if dp[i][j] > maxEdge {
					maxEdge = dp[i][j]
				}
			}
		}
	}

	return maxEdge * maxEdge
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}
