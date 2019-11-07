func getRow(n int) []int {
	res := make([]int, 1, n+1)
	res[0] = 1
	if n == 0 {
		return res
	}

	for i := 0; i < n; i++ {
		res = append(res, 1)
		for j := len(res) - 2; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
