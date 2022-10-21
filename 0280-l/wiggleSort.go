package main

import (
	"fmt"
	"sort"
)

func wiggleSortSlow(nums []int) {
	sort.Ints(nums)
	for i := 1; i < len(nums)-1; i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}

// 进阶版: O(N) 复杂度
func wiggleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		// 阶段 [a < b] c
		// 如果不 swap, [a < b] 不被破坏
		// 如果 swap 了, 也必然满足 [a < c] > b
		if (i%2 == 0 && nums[i] > nums[i+1]) ||
			(i%2 == 1 && nums[i] < nums[i+1]) {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}

func main() {
	nums := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
}
