func permuteUnique(nums []int) [][]int {
	permutation := make([]int, len(nums))
	numInPermutation := make([]bool, len(nums))

	var permutations [][]int
	makePermutation(0, nums, permutation, numInPermutation, &permutations)
	return permutations
}

func makePermutation(idx int, nums, permutation []int, numInPermutation []bool, permutations *[][]int) {
	if idx == len(nums) {
		deepcopy := make([]int, len(nums))
		copy(deepcopy, permutation)
		*permutations = append(*permutations, deepcopy)
		return
	}

	numInPermutationAtIdx := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if !numInPermutation[i] && !numInPermutationAtIdx[nums[i]] {
			numInPermutation[i] = true
			numInPermutationAtIdx[nums[i]] = true

			permutation[idx] = nums[i]
			makePermutation(idx+1, nums, permutation, numInPermutation, permutations)

			numInPermutation[i] = false
		}
	}
}
