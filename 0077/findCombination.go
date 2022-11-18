package main

import "fmt"

// dfs
func combine(n int, k int) [][]int {
	var combinations [][]int
	combineHelper(n, k, 1, []int{}, &combinations)
	return combinations
}

func combineHelper(n, k, cur int, combination []int, combinations *[][]int) {
	if len(combination) == k {
		deepcopy := make([]int, k)
		copy(deepcopy, combination)
		*combinations = append(*combinations, deepcopy)
		return
	}

	if len(combination) + (n-cur+1) < k {
		return
	}
	
	combineHelper(n, k, cur+1, append(combination, cur), combinations)
	combineHelper(n, k, cur+1, combination, combinations)
}

func main() {
	fmt.Println(combine(5, 2))
}
