package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	doFindSubsets(nums, []int{}, &result)
	return result
}

func doFindSubsets(nums []int, subset []int, result *[][]int) {
	if len(nums) == 0 {
		deepcopy := make([]int, len(subset))
		copy(deepcopy, subset)
		*result = append(*result, deepcopy)
		return
	}

	doFindSubsets(nums[1:], subset, result)
	doFindSubsets(nums[1:], append(subset, nums[0]), result)
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
