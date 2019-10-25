import "sort"

func combinationSum2(candidates []int, target int) [][]int {
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

	findCombination(candidates[1:], append(solution, candidates[0]), target-candidates[0], result)
	var i int
	for i = 1; i < len(candidates) && candidates[i] == candidates[0]; i++ {
	}
	findCombination(candidates[i:], solution, target, result)
}

