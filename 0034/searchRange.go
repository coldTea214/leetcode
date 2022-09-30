func searchRange(nums []int, target int) []int {
	// 查看 target 是否存在
	index := search(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}

	// 每次能排除一半的数，即便数组全是 target，也是logn复杂度
	// 查找第一个
	first := index
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}
		first = f
	}

	// 查找最后一个
	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}
		// 细节处理
		last += l + 1
	}

	return []int{first, last}
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
