package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
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

	doFindCombination(candidates[1:], target-candidates[0], append(solution, candidates[0]), solutions)
	// [1 1 1 2 3], 5
	// 1 1 1 2
	//     3
	// 2 3 
	var i int
	for i = 1; i < len(candidates) && candidates[i] == candidates[0]; i++ {
	}
	doFindCombination(candidates[i:], target, solution, solutions)
}

func main() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
