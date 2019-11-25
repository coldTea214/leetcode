func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	return selection(nums, 0, len(nums)-1, len(nums)-k)
}

func selection(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	p := partition(nums, l, r)
	if k == p {
		return nums[p]
	} else if k < p {
		return selection(nums, l, p-1, k)
	} else {
		return selection(nums, p+1, r, k)
	}
}

func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
