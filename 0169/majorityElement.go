func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			count++
		} else if count > 0 {
			count--
		} else {
			res, count = nums[i], 1
		}
	}
	return res
}
