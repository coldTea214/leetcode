package main

import "fmt"

func permute(nums []int) [][]int {
	numInPermutation := make(map[int]bool)
	var permutations [][]int
	doPermute(nums, numInPermutation, []int{}, &permutations)
	return permutations
}

func doPermute(nums []int, numInPermutation map[int]bool, permutation []int, permutations *[][]int) {
	if len(permutation) == len(nums) {
		deepcopy := make([]int, len(nums))
		copy(deepcopy, permutation)
		*permutations = append(*permutations, deepcopy)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !numInPermutation[nums[i]] {
			numInPermutation[nums[i]] = true
			doPermute(nums, numInPermutation, append(permutation, nums[i]), permutations)
			numInPermutation[nums[i]] = false
		}
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
