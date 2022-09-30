package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	numInPermutation := make(map[int]bool)
	var permutations [][]int
	doPermuteUnique(nums, numInPermutation, []int{}, &permutations)
	return permutations
}

func doPermuteUnique(nums []int, numInPermutation map[int]bool, permutation []int, permutations *[][]int) {
	if len(permutation) == len(nums) {
		deepcopy := make([]int, len(nums))
		copy(deepcopy, permutation)
		*permutations = append(*permutations, deepcopy)
		return
	}

	numInPermutationThisLevel := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if !numInPermutation[nums[i]] && !numInPermutationThisLevel[nums[i]] {
			numInPermutation[nums[i]] = true
			numInPermutationThisLevel[nums[i]] = true
			doPermuteUnique(nums, numInPermutation, append(permutation, nums[i]), permutations)
			numInPermutation[nums[i]] = false
		}
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
