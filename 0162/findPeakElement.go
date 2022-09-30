// 因为两端是极低值，只要一路跟着大的值走，必然有一个峰
func findPeakElement(nums []int) int {
	low, high := 0, len(nums)
	for low < high {
		mid = (high + low) >> 1
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
