func generate(n int) [][]int {
	res := make([][]int, n)
	for i := 1; i <= n; i++ {
		res[i-1] = make([]int, i)
		res[i-1][0] = 1
		res[i-1][i-1] = 1
	}

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
