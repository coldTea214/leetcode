// 前置题 0026
func removeDuplicates(nums []int) int {
	slow := 2
	for quick := slow; quick < len(nums); quick++ {
		if nums[quick] == nums[slow-2] {
			continue
		}
		nums[slow] = nums[quick]
		slow++
	}
	return slow
}
