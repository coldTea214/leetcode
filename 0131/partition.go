package main

import "fmt"

// dfs
func partition(s string) [][]string {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := range s {
		isPalindrome[i] = make([]bool, n)
		for j := range s {
			isPalindrome[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i+1][j-1]
		}
	}

	solutions := [][]string{}
	partitionHelper(0, s, isPalindrome, []string{}, &solutions)
	return solutions
}

func partitionHelper(idx int, s string, isPalindrome [][]bool, solution []string, solutions *[][]string) {
	if idx == len(s) {
		deepcopy := make([]string, len(solution))
		copy(deepcopy, solution)
		*solutions = append(*solutions, deepcopy)
		return
	}

	for i := idx; i < len(s); i++ {
		if isPalindrome[idx][i] {
			partitionHelper(i+1, s, isPalindrome, append(solution, s[idx:i+1]), solutions)
		}
	}
}

func main() {
	fmt.Println(partition("aab"))
}
