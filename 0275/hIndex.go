func hIndex(citations []int) int {
	count := len(citations)
	low, high := 0, count-1
	var mid int
	for low <= high {
		mid = (low + high) >> 1
		if citations[mid] == count-mid {
			return citations[mid]
		} else if citations[mid] > count-mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return len(citations) - low
}
