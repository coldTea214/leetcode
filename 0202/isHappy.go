func isHappy(n int) bool {
	numAppeared := make(map[int]bool)
	for !numAppeared[n] {
		numAppeared[n] = true
		n = transform(n)
	}
	return n == 1
}

func transform(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
