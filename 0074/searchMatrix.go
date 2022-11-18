// 后置题 0240
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
	r, i, j := 0, 0, m-1
	for i <= j {
		r = (i + j) / 2
		if matrix[r][n-1] < target {
			i = r + 1
		} else if matrix[r][0] > target {
			j = r - 1
		} else {
			break
		}
	}

	c, i, j := 0, 0, n-1
	for i <= j {
		c = (i + j) / 2
		if matrix[r][c] < target {
			i = c + 1
		} else if target < matrix[r][c] {
			j = c - 1
		} else {
			return true
		}
	}
	return false
}
