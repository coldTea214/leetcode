package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	solutions := [][]int{}
	doFindCombination(candidates, target, []int{}, &solutions)
	return solutions
}

func doFindCombination(candidates []int, target int, solution []int, solutions *[][]int) {
	if target == 0 {
		deepcopy := make([]int, len(solution))
		copy(deepcopy, solution)
		*solutions = append(*solutions, deepcopy)
		return
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	doFindCombination(candidates, target-candidates[0], append(solution, candidates[0]), solutions)
	doFindCombination(candidates[1:], target, solution, solutions)
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
