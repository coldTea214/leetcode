func combinationSum3(k int, n int) [][]int {
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	res := [][]int{}
	solution := []int{}
	findCombination(candidates, solution, k, n, &res)

	return res
}

func findCombination(candidates, solution []int, k, n int, result *[][]int) {
	if n == 0 && len(solution) == k {
		deepcopy := make([]int, len(solution))
		copy(deepcopy, solution)
		*result = append(*result, deepcopy)
	}

	if len(candidates) == 0 || n < candidates[0] || len(solution) == k {
		return
	}

	findCombination(candidates[1:], append(solution, candidates[0]), k, n-candidates[0], result)
	findCombination(candidates[1:], solution, k, n, result)
}
