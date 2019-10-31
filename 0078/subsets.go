func subsets(nums []int) [][]int {
	var result [][]int
	var subset []int
	findSubsets(nums, 0, subset, &result)
	return result
}

func findSubsets(nums []int, count int, subset []int, result *[][]int) {
	if count == len(nums) {
		deepcopy := make([]int, len(subset))
		copy(deepcopy, subset)
		*result = append(*result, deepcopy)
		return
	}

	findSubsets(nums, count+1, subset, result)
	findSubsets(nums, count+1, append(subset, nums[count]), result)
}
