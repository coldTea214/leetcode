func combine(n int, k int) [][]int {
	combination := make([]int, k)
	var combinations [][]int
	findCombination(n, k, 0, 1, combination, &combinations)
	return combinations
}

func findCombination(n, k, idx, begin int, combination []int, combinations *[][]int) {
	if idx == k {
		deepcopy := make([]int, k)
		copy(deepcopy, combination)
		*combinations = append(*combinations, deepcopy)
		return
	}

	for i := begin; i <= n+1-k+idx; i++ {
		combination[idx] = i
		findCombination(n, k, idx+1, i+1, combination, combinations)
	}
}
