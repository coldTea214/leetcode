func jump(nums []int) int {
	minStep := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			loc := i + j
			if loc >= len(nums) {
				break
			}
			if minStep[loc] == 0 || minStep[i]+1 < minStep[loc] {
				minStep[loc] = minStep[i] + 1
			}
		}
	}
	return minStep[len(nums)-1]
}
