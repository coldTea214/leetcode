// 前置题 0033
// 变化点是 nums 可能有相同数字
func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if target == nums[mid] {
			return true
		} else if nums[mid] < nums[high] {
			// 4 5 1 2 3
			//     m
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else if nums[mid] > nums[high] {
			// 3 4 5 1 2
			//     m
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			high--
		}
	}
	return false
}
