import "strconv"

func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))
	for _, s := range tokens {
		switch s {
		case "+", "-", "*", "/":
			b, a := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, compute(a, b, s))
		default:
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
	}
	return nums[0]
}

func compute(a, b int, opt string) int {
	switch opt {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
