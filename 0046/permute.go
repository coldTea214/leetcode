func permute(nums []int) [][]int {
	permutation := make([]int, len(nums))
	numInPermutation := make([]bool, len(nums))

	var permutations [][]int
	makePermutation(0, nums, permutation, numInPermutation, &permutations)
	return permutations
}

func makePermutation(count int, nums, permutation []int, numInPermutation []bool, permutations *[][]int) {
	if count == len(nums) {
		deepcopy := make([]int, len(nums))
		copy(deepcopy, permutation)
		*permutations = append(*permutations, deepcopy)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !numInPermutation[i] {
			numInPermutation[i] = true

			permutation[count] = nums[i]
			makePermutation(count+1, nums, permutation, numInPermutation, permutations)

			numInPermutation[i] = false
		}
	}
}
