func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 0
	for quick := 1; quick < len(nums); quick++ {
		if nums[slow] == nums[quick] {
			continue
		}
		slow++
		nums[slow], nums[quick] = nums[quick], nums[slow]
	}
	return slow + 1
}
