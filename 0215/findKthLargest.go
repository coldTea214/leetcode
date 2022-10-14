package main

import (
	"math/rand"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	p := randomPartition(nums, l, r)
	if k == p {
		return nums[p]
	} else if k < p {
		return quickSelect(nums, l, p-1, k)
	} else {
		return quickSelect(nums, p+1, r, k)
	}
}

// 常规快排最坏时间复杂度是O(n^2)
// 引入随机化期望时间复杂度是O(n)，算法导论中有证明
func randomPartition(nums []int, l, r int) int {
	i := rand.Int()%(r-l+1) + 1
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
