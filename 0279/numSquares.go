import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = i
	}

	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
		dp[i*i] = 1
	}

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				dp[i] = dp[i-p] + 1
			}
		}
	}
	return dp[n]
}
