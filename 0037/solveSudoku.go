// dfs
func solveSudoku(board [][]byte) {
	var appearInRow, appearInCol, appearInBlock [9][10]bool
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '0'
			appearInRow[i][num] = true
			appearInCol[j][num] = true
			appearInBlock[i/3*3+j/3][num] = true
		}
	}

	solveSudokuHelper(0, appearInRow, appearInCol, appearInBlock, board)
}

// dfs 仅求一个解时, helper 函数 return bool
func solveSudokuHelper(count int, appearInRow, appearInCol, appearInBlock [9][10]bool, board [][]byte) bool {
	if count == 81 {
		return true
	}

	row, col := count/9, count%9
	if board[row][col] != '.' {
		return solveSudokuHelper(count+1, appearInRow, appearInCol, appearInBlock, board)
	}

	for b := byte('1'); b <= '9'; b++ {
		num := b - '0'
		if appearInRow[row][num] || appearInCol[col][num] || appearInBlock[row/3*3+col/3][num] {
			continue
		}

		board[row][col] = b
		appearInRow[row][num], appearInCol[col][num], appearInBlock[row/3*3+col/3][num] = true, true, true
		if solveSudokuHelper(count+1, appearInRow, appearInCol, appearInBlock, board) {
			return true
		}
		appearInRow[row][num], appearInCol[col][num], appearInBlock[row/3*3+col/3][num] = false, false, false
	}
	board[row][col] = '.'

	return false
}
