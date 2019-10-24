func removeElement(nums []int, val int) int {
	// i指向第一个值为val的位置
	// right指向最后一个不为val的位置
	left, right := 0, len(nums)-1
	for {
		for left < len(nums) && nums[left] != val {
			left++
		}

		for right >= 0 && nums[right] == val {
			right--
		}

		if left >= right {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	return left
}
