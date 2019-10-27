func solveNQueens(n int) [][]string {
	// appearInLeftSlash，皇后位于左斜线（例：[0,0]->[7,7])
	// 以 5 皇后为例
	// 第1条：[4,0]
	// 第2条：[3,0] [4,1]
	// 第3条：[2,0] [3,1] [4,2]
	// ...
	// 第5条：[0,0] [1,1] [2,2] [3,3] [4,4]
	// 第6条：[0,1] [1,2] [2,3] [3,4]
	// ...
	// [i,j] 位于第 5-(i-j) 条左斜线
	// 右斜线：
	// 第1条：[0,0]
	// 第2条：[0,1] [1,0]
	// 第3条：[0,2] [1,1] [2,0]
	// ...
	// 第5条：[0,4] [1,3] [2,2] [3,1] [4,0]
	// 第6条：[1,4] [2,3] [3,2] [4,1]
	// ...
	// [i,j] 位于第 (i+j)+1 条右斜线
	appearInCol := make([]bool, n)
	appearInLeftSlash, appearInRightSlash := make([]bool, n*2-1), make([]bool, n*2-1)

	solution := make([]string, n)
	var solutions [][]string
	doSolve(0, solution, &solutions, appearInCol, appearInLeftSlash, appearInRightSlash)
	return solutions
}

func doSolve(count int, solution []string, solutions *[][]string, appearInCol, appearInLeftSlash, appearInRightSlash []bool) {
	i, n := count, len(appearInCol)
	if count == n {
		deepcopy := make([]string, n)
		copy(deepcopy, solution)
		*solutions = append(*solutions, deepcopy)
	}

	rowInSolution := make([]byte, n)
	for j := 0; j < n; j++ {
		rowInSolution[j] = '.'
	}
	for j := 0; j < n; j++ {
		if !appearInCol[j] && !appearInLeftSlash[n-i+j-1] && !appearInRightSlash[i+j] {
			appearInCol[j], appearInLeftSlash[n-i+j-1], appearInRightSlash[i+j] = true, true, true
			rowInSolution[j] = 'Q'

			solution[count] = string(rowInSolution)
			doSolve(count+1, solution, solutions, appearInCol, appearInLeftSlash, appearInRightSlash)

			rowInSolution[j] = '.'
			appearInCol[j], appearInLeftSlash[n-i+j-1], appearInRightSlash[i+j] = false, false, false
		}
	}
}
