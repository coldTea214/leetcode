func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}

	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}
