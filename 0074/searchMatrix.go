func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	if target < matrix[0][0] || matrix[m-1][n-1] < target {
		return false
	}

	// 寻找行
	foundRow := false
	r, i, j := 0, 0, m-1
	for i <= j && !foundRow {
		r = (i + j) / 2
		switch {
		case matrix[r][n-1] < target:
			i = r + 1
		case matrix[r][0] > target:
			j = r - 1
		default:
			foundRow = true
		}
	}

	c, i, j := 0, 0, n-1
	for i <= j {
		c = (i + j) / 2
		switch {
		case matrix[r][c] < target:
			i = c + 1
		case target < matrix[r][c]:
			j = c - 1
		default:
			return true
		}
	}

	return matrix[r][c] == target
}
