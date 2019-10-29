func mySqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := (low + high) >> 1
		tmp := mid * mid
		if tmp == x {
			return mid
		} else if tmp < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}
