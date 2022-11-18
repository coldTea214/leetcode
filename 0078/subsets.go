package main

import "fmt"

// dfs
// 后置题 0090
func subsets(nums []int) [][]int {
	var result [][]int
	subsetsHelper(nums, []int{}, &result)
	return result
}

func subsetsHelper(nums []int, subset []int, result *[][]int) {
	if len(nums) == 0 {
		deepcopy := make([]int, len(subset))
		copy(deepcopy, subset)
		*result = append(*result, deepcopy)
		return
	}

	subsetsHelper(nums[1:], subset, result)
	subsetsHelper(nums[1:], append(subset, nums[0]), result)
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
