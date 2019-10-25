import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	solution := []int{}
	findCombination(candidates, solution, target, &res)

	return res
}

func findCombination(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		deepcopy := make([]int, len(solution))
		copy(deepcopy, solution)
		*result = append(*result, deepcopy)
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	findCombination(candidates, append(solution, candidates[0]), target-candidates[0], result)
	findCombination(candidates[1:], solution, target, result)
}
