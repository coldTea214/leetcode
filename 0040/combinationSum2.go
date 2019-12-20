import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	solutions := [][]int{}

	var findCombination func([]int, []int, int)
	findCombination = func(candidates, solution []int, target int) {
		if target == 0 {
			deepcopy := make([]int, len(solution))
			copy(deepcopy, solution)
			solutions = append(solutions, deepcopy)
		}

		if len(candidates) == 0 || target < candidates[0] {
			return
		}

		findCombination(candidates[1:], append(solution, candidates[0]), target-candidates[0])
		var i int
		for i = 1; i < len(candidates) && candidates[i] == candidates[0]; i++ {
		}
		findCombination(candidates[i:], solution, target)
	}
	findCombination(candidates, []int{}, target)

	return solutions
}
