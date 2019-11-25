func containsDuplicate(nums []int) bool {
	exist := make(map[int]bool, len(nums))
	for _, n := range nums {
		if exist[n] {
			return true
		}
		exist[n] = true
	}
	return false
}
