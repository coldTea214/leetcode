import "sort"

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	var subset []int
	sort.Ints(nums)
	findSubsets(nums, 0, subset, &result)
	return result
}

func findSubsets(nums []int, index int, subset []int, result *[][]int) {
	deepcopy := make([]int, len(subset))
	copy(deepcopy, subset)
	*result = append(*result, deepcopy)

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		findSubsets(nums, i+1, append(subset, nums[i]), result)
	}
}
