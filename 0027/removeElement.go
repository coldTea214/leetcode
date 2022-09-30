func removeElement(nums []int, val int) int {
	// left指向第一个值为val的位置
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

// 如果不允许改变其他元素顺序
func removeElement2(nums []int, val int) int {
	slow := 0
	for quick := 0; quick < len(nums); quick++ {
		if nums[quick] == val {
			continue
		}
		nums[slow] = nums[quick]
		slow++
	}
	return slow
}
