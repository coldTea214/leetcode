// 前置题 0198
func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}

	return max(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
