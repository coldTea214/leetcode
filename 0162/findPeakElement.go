func findPeakElement(nums []int) int {
	low, high := -1, len(nums)
	mid := (low + high) >> 1

	for low+1 <= mid && mid < high-1 {
		if nums[mid] < nums[mid+1] {
			low = mid
		} else {
			high = mid + 1
		}
		mid = (high + low) >> 1
	}

	return low + 1
}
