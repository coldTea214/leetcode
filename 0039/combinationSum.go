import "sort"

func combinationSum(candidates []int, target int) [][]int {
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

		findCombination(candidates, append(solution, candidates[0]), target-candidates[0])
		findCombination(candidates[1:], solution, target)
	}
	findCombination(candidates, []int{}, target)

	return solutions
}
