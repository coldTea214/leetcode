func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := 1

	/**
	 * 利用 matrix[i][0] = 0 表示，第 i 行中含有 0
	 * 利用 matrix[0][j] = 0 表示，第 j 列中含有 0
	 * 特别地，
	 * matrix[0][0] = 0 仅表示，第 0 行中含有 0
	 * Col0 = 0 表示，第 0 列中含有 0
	 */
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	/**
	 * 从下往上，从右往左 扫描矩阵
	 * 并根据前面的标记修改 matrix[i][j] 的值
	 * NOTICE: 循环的顺序很重要
	 * 需要保证 matrix[i][0] 和 matrix[0][j] 被修改后，不再做为别的修改的标记
	 * 要不然的话，标记有可能会被污染
	 */
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 == 0 {
			matrix[i][0] = 0
		}
	}
}
