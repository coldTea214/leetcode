func maxSlidingWindow(nums []int, k int) []int {
	if k <= 1 {
		return nums
	}

	nLen := len(nums)
	leftMax := make([]int, nLen)
	for i := 0; i < nLen; i++ {
		if i%k == 0 {
			leftMax[i] = nums[i]
		} else {
			leftMax[i] = max(nums[i], leftMax[i-1])
		}
	}

	rightMax := make([]int, nLen)
	rightMax[nLen-1] = nums[nLen-1]
	for j := nLen - 2; j >= 0; j-- {
		if j%k == 0 {
			rightMax[j] = nums[j]
		} else {
			rightMax[j] = max(nums[j], rightMax[j+1])
		}
	}

	res := make([]int, nLen-k+1)
	for i := 0; i <= nLen-k; i++ {
		res[i] = max(rightMax[i], leftMax[i+k-1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
