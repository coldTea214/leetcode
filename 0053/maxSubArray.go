func maxSubArray(nums []int) int {
	sum, maxSum := max(0, nums[0]), nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		sum = max(0, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
