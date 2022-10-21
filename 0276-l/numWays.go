func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	a, b, c := k, k*k, 0
	for i := 2; i < n; i++ {
		// 3个桩时
		// 不考虑重色则是 k*k*k
		// 去掉不满足条件的 k*k
		c = (a + b) * (k - 1)
		a, b = b, c
	}
	return b
}
