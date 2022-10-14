package main

import (
	"fmt"
)

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	bucket := map[int]int{}
	for i, num := range nums {
		id := bucketID(num, t+1)
		if _, ok := bucket[id]; ok {
			return true
		}
		if n, ok := bucket[id-1]; ok && num-n <= t {
			return true
		}
		if n, ok := bucket[id+1]; ok && n-num <= t {
			return true
		}
		bucket[id] = num
		if i >= k {
			delete(bucket, bucketID(nums[i-k], t+1))
		}
	}
	return false
}

func bucketID(num, width int) int {
	if num >= 0 {
		return num / width
	}
	return (num+1)/width - 1
}

func main() {
	// fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 1, 1}, 1, 0))
}
