func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	digits[len(digits)-1]++
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	if digits[0] == 10 {
		digits = append([]int{1, 0}, digits[1:]...)
	}
	return digits
}
