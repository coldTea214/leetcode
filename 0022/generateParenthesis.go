package main

import "fmt"

func generateParenthesis(n int) []string {
	var solutions []string
	doGenerate(n, 0, 0, "", &solutions)
	return solutions
}

func doGenerate(n, lNum, rNum int, solution string, solutions *[]string) {
	if lNum == n && rNum == n {
		*solutions = append(*solutions, solution)
		return
	}

	if lNum < n {
		doGenerate(n, lNum+1, rNum, solution+"(", solutions)
	}
	if rNum < lNum {
		doGenerate(n, lNum, rNum+1, solution+")", solutions)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
