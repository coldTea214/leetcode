func partition(s string) [][]string {
	solutions := [][]string{}
	solution := make([]string, 0, len(s))
	findPartition(s, 0, solution, &solutions)
	return solutions
}

func findPartition(s string, i int, cur []string, solutions *[][]string) {
	if i == len(s) {
		deepcopy := make([]string, len(cur))
		copy(deepcopy, cur)
		*solutions = append(*solutions, deepcopy)
		return
	}

	for j := i; j < len(s); j++ {
		if isPalindrome(s[i : j+1]) {
			findPartition(s, j+1, append(cur, s[i:j+1]), solutions)
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
