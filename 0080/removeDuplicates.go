func removeDuplicates(nums []int) int {
	size := len(nums)

	slow := 2
	for quick := slow; quick < size; quick++ {
		if nums[slow-2] != nums[quick] {
			nums[slow] = nums[quick]
			slow++
		}
	}
	return slow
}
