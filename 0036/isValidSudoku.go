func isValidSudoku(board [][]byte) bool {
	appearInRow := [10][10]bool{}
	appearInCol := [10][10]bool{}
	appearInBlock := [10][10]bool{}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0'
			if appearInRow[i][num] || appearInCol[j][num] || appearInBlock[i/3*3+j/3][num] {
				return false
			}

			appearInRow[i][num] = true
			appearInCol[j][num] = true
			appearInBlock[i/3*3+j/3][num] = true
		}
	}
	return true
}
