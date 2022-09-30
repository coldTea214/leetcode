package main

import "fmt"

func partition(s string) [][]string {
	solutions := [][]string{}
	doPartition(s, []string{}, &solutions)
	return solutions
}

func doPartition(s string, solution []string, solutions *[][]string) {
	if len(s) == 0 {
		deepcopy := make([]string, len(solution))
		copy(deepcopy, solution)
		*solutions = append(*solutions, deepcopy)
		return
	}

	for i := 0; i < len(s); i++ {
		if isPalindrome(s[:i+1]) {
			doPartition(s[i+1:], append(solution, s[:i+1]), solutions)
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
