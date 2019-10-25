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

	doSolve(board, 0, appearInRow, appearInCol, appearInBlock)
}

func doSolve(board [][]byte, count int, appearInRow, appearInCol, appearInBlock [9][10]bool) bool {
	if count == 81 {
		return true
	}

	row, col := count/9, count%9
	if board[row][col] != '.' {
		return doSolve(board, count+1, appearInRow, appearInCol, appearInBlock)
	}

	for b := byte('1'); b <= '9'; b++ {
		num := b - '0'
		if !appearInRow[row][num] && !appearInCol[col][num] && !appearInBlock[row/3*3+col/3][num] {
			board[row][col] = b
			appearInRow[row][num], appearInCol[col][num], appearInBlock[row/3*3+col/3][num] = true, true, true
			if doSolve(board, count+1, appearInRow, appearInCol, appearInBlock) {
				return true
			}
			appearInRow[row][num], appearInCol[col][num], appearInBlock[row/3*3+col/3][num] = false, false, false
		}
	}

	board[row][col] = '.'

	return false
}
