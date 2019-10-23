import "math"

// 执行用时 : 0 ms , 在所有 golang 提交中击败了 100.00% 的用户
// 内存消耗 : 2.2 MB , 在所有 golang 提交中击败了 90.87% 的用户

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
