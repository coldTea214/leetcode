package main

import "fmt"

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	// 把 k 移到 nums[k-1]
	for i := 0; i < len(nums); {
		k := nums[i]
		if k <= 0 || k > len(nums) || nums[k-1] == k {
			i++
			continue
		}
		nums[i], nums[k-1] = nums[k-1], nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
