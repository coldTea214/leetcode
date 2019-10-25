func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	// 把 k 移到 nums[k-1]
	for i := 0; i < len(nums); {
		k := nums[i]
		if k <= 0 || k > len(nums) || nums[k-1] == k {
			i++
			continue
		}
		nums[i], nums[k-1] = nums[k-1], nums[i]
		if nums[i] == i+1 {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
