// dfs
func solve(board [][]byte) {
	// 从边界的 O 出发，能抵达的 O 更改为 Y
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
				if board[i][j] == 'O' {
					solveHelper(i, j, board)
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solveHelper(i, j int, board [][]byte) {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = 'Y'

		solveHelper(i-1, j, board)
		solveHelper(i, j-1, board)
		solveHelper(i+1, j, board)
		solveHelper(i, j+1, board)
	}
}
