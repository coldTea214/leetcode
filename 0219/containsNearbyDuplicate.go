func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	appearLoc := make(map[int]int, len(nums))
	for i, n := range nums {
		loc, ok := appearLoc[n]
		if ok && i-loc <= k {
			return true
		}
		appearLoc[n] = i
	}
	return false
}
