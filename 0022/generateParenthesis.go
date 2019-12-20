func generateParenthesis(n int) []string {
	var solutions []string
	var doGenerate func(int, int, int, string)
	doGenerate = func(n, lNum, rNum int, solution string) {
		if lNum == n && rNum == n {
			solutions = append(solutions, solution)
			return
		}

		if lNum < n {
			doGenerate(n, lNum+1, rNum, solution+"(")
		}
		if rNum < lNum {
			doGenerate(n, lNum, rNum+1, solution+")")
		}
	}
	doGenerate(n, 0, 0, "")
	return solutions
}
