// dfs
func totalNQueens(n int) int {
	appearInCol := make([]bool, n)
	appearInLeftSlash, appearInRightSlash := make([]bool, n*2-1), make([]bool, n*2-1)

	var total int
	totalNQueensHelper(0, appearInCol, appearInLeftSlash, appearInRightSlash, &total)
	return total
}

func totalNQueensHelper(count int, appearInCol, appearInLeftSlash, appearInRightSlash []bool, total *int) {
	i, n := count, len(appearInCol)
	if i == n {
		*total++
		return
	}

	for j := 0; j < n; j++ {
		if !appearInCol[j] && !appearInLeftSlash[n-i+j-1] && !appearInRightSlash[i+j] {
			appearInCol[j], appearInLeftSlash[n-i+j-1], appearInRightSlash[i+j] = true, true, true

			totalNQueensHelper(count+1, appearInCol, appearInLeftSlash, appearInRightSlash, total)

			appearInCol[j], appearInLeftSlash[n-i+j-1], appearInRightSlash[i+j] = false, false, false
		}
	}
}
