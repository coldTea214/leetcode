package main

import (
	"fmt"
	"sort"
)

// dfs
func permuteUnique(nums []int) [][]int {
	numInPermutation := make(map[int]bool)
	var permutations [][]int
	sort.Ints(nums)
	permuteUniqueHelper(nums, numInPermutation, []int{}, &permutations)
	return permutations
}

func permuteUniqueHelper(nums []int, numInPermutation map[int]bool, permutation []int, permutations *[][]int) {
	if len(permutation) == len(nums) {
		deepcopy := make([]int, len(nums))
		copy(deepcopy, permutation)
		*permutations = append(*permutations, deepcopy)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 此时只能以索引i来唯一确认num出现在了排列里，而非nums[i]
		if !numInPermutation[i] {
			numInPermutation[i] = true
			permuteUniqueHelper(nums, numInPermutation, append(permutation, nums[i]), permutations)
			numInPermutation[i] = false
			for ; i+1 < len(nums) && nums[i+1] == nums[i]; i++ {
			}
		}
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
