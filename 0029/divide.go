import "math"

func divide(dividend int, divisor int) int {
	var resNegtive bool
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		resNegtive = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	var divisorMultiples []int
	multiple := divisor
	for multiple <= dividend {
		divisorMultiples = append(divisorMultiples, multiple)
		multiple += multiple
	}

	var res int
	for i := len(divisorMultiples) - 1; i >= 0; i-- {
		if dividend >= divisorMultiples[i] {
			dividend -= divisorMultiples[i]
			res += 1 << uint(i)
		}
	}

	if resNegtive {
		res = -res
	}
	if res > math.MaxInt32 {
		res = math.MaxInt32
	}
	return res
}
