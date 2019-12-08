func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	i, j := m-1, 0
	for 0 <= i && j < n {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return false
}
