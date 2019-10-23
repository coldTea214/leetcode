import "math"

func reverse(x int) int {
	var isNegative bool
	if x < 0 {
		isNegative = true
		x = -x
	}

	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	if result > math.MaxInt32 {
		return 0
	}
	if isNegative {
		result = -result
	}
	return result
}
