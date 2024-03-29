// 前置题 0040
func combinationSum3(k int, n int) [][]int {
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	solutions := [][]int{}
	combinationSumHelper(candidates, k, n, []int{}, &solutions)
	return solutions
}

func combinationSumHelper(candidates []int, k, n int, solution []int, solutions *[][]int) {
	if n == 0 && len(solution) == k {
		deepcopy := make([]int, len(solution))
		copy(deepcopy, solution)
		*solutions = append(*solutions, deepcopy)
		return
	}

	if len(candidates) == 0 || n < candidates[0] || len(solution) == k {
		return
	}

	combinationSumHelper(candidates[1:], k, n-candidates[0], append(solution, candidates[0]), solutions)
	combinationSumHelper(candidates[1:], k, n, solution, solutions)
}
