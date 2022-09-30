import "math"

/*
-2 -3  3
-5 -10 1
10 30  -5
对应的 dp 为：
7   5   2 max
6   11  5 max
1   1   6 1
max max 1 max
*/
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 1
	}
	n := len(dungeon[0])
	if n == 0 {
		return 1
	}

	// dp[i][j] 到达 (i,j) 前的健康值
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	health := 0
	dp[m][n-1], dp[m-1][n] = 1, 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			health = min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			dp[i][j] = max(health, 1)
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
