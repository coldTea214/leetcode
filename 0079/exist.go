func exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0 {
		return false
	}

	col := len(board[0])
	if col == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}

	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(i, j, row, col, 0, word, board, visited) {
				return true
			}
		}
	}

	return false
}

func dfs(curRow, curCol, row, col, count int, word string, board [][]byte, visited [][]bool) bool {
	if count == len(word) {
		return true
	}

	if curRow < 0 || curRow >= row || curCol < 0 || curCol >= col ||
		visited[curRow][curCol] ||
		board[curRow][curCol] != word[count] {
		return false
	}

	visited[curRow][curCol] = true

	if dfs(curRow-1, curCol, row, col, count+1, word, board, visited) ||
		dfs(curRow+1, curCol, row, col, count+1, word, board, visited) ||
		dfs(curRow, curCol-1, row, col, count+1, word, board, visited) ||
		dfs(curRow, curCol+1, row, col, count+1, word, board, visited) {
		return true
	}

	visited[curRow][curCol] = false

	return false
}
