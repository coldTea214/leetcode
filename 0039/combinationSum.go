package main

import (
	"fmt"
	"sort"
)

// dfs
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	solutions := [][]int{}
	combinationSumHelper(candidates, target, []int{}, &solutions)
	return solutions
}

func combinationSumHelper(candidates []int, target int, solution []int, solutions *[][]int) {
	if target == 0 {
		deepcopy := make([]int, len(solution))
		copy(deepcopy, solution)
		*solutions = append(*solutions, deepcopy)
		return
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	combinationSumHelper(candidates, target-candidates[0], append(solution, candidates[0]), solutions)
	combinationSumHelper(candidates[1:], target, solution, solutions)
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
