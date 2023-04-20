// 后置题 0152
func maxSubArray(nums []int) int {
	curMax, maxSum := max(0, nums[0]), nums[0]
	for i := 1; i < len(nums); i++ {
		curMax += nums[i]
		if curMax > maxSum {
			maxSum = curMax
		}
		curMax = max(0, curMax)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
