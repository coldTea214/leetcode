package main

import "fmt"

// dfs
func generateParenthesis(n int) []string {
	var solutions []string
	generateParenthesisHelper(n, 0, 0, "", &solutions)
	return solutions
}

func generateParenthesisHelper(n, lNum, rNum int, solution string, solutions *[]string) {
	if lNum == n && rNum == n {
		*solutions = append(*solutions, solution)
		return
	}

	if lNum < n {
		generateParenthesisHelper(n, lNum+1, rNum, solution+"(", solutions)
	}
	if rNum < lNum {
		generateParenthesisHelper(n, lNum, rNum+1, solution+")", solutions)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
