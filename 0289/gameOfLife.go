func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}

	check := func(i, j int) {
		// 检查 board[i][j] 一圈有几个活细胞
		liveNum := 0
		for r := i - 1; r <= i+1; r++ {
			for c := j - 1; c <= j+1; c++ {
				if 0 <= r && r < m && 0 <= c && c < n &&
					(r != i || c != j) &&
					(board[r][c] == 1 || board[r][c] == 2) {
					liveNum++
				}
			}
		}

		if board[i][j] == 1 && (liveNum < 2 || liveNum > 3) {
			// 规则1、3
			board[i][j] = 2
		} else if board[i][j] == 0 && liveNum == 3 {
			// 规则4
			board[i][j] = 3
		}
		// 规则2
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			check(i, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] %= 2
		}
	}
}
