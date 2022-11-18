package main

import (
	"fmt"
	"sort"
)

// 前置题 0078
// 变化点是 nums可以重复
func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	subsetsWithDupHelper(nums, []int{}, &result)
	return result
}

func subsetsWithDupHelper(nums []int, subset []int, result *[][]int) {
	if len(nums) == 0 {
		deepcopy := make([]int, len(subset))
		copy(deepcopy, subset)
		*result = append(*result, deepcopy)
		return
	}

	subsetsWithDupHelper(nums[1:], append(subset, nums[0]), result)
	var i int
	for i = 1; i < len(nums) && nums[i] == nums[0]; i++ {
	}
	subsetsWithDupHelper(nums[i:], subset, result)
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
