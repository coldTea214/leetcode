func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	aInt := toIntSlice(a, len(a))
	bInt := toIntSlice(b, len(a))

	return makeString(add(aInt, bInt))
}

// [1,1], [0, 1]
func toIntSlice(s string, maxLen int) []int {
	res := make([]int, maxLen)
	ls := len(s)
	for i, b := range s {
		res[maxLen-ls+i] = int(b - '0')
	}
	return res
}

func add(a, b []int) []int {
	var sum, carry int
	res := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		sum = a[i] + b[i] + carry
		res[i] = sum % 2
		carry = sum / 2
	}

	if carry == 1 {
		return append([]int{1}, res...)
	}
	return res
}

func makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}
	return string(bytes)
}
