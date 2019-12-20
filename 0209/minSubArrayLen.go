func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	l, r := 0, 0
	minLen := n + 1
	sum := 0
	for l < n {
		if sum < s {
			if r == n {
				break
			}
			sum += nums[r]
			r++
		} else {
			if r-l < minLen {
				minLen = r - l
			}
			sum -= nums[l]
			l++
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}
