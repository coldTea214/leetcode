package main

import "fmt"

// 后置题 0080
func removeDuplicates(nums []int) int {
	slow := 1
	for quick := slow; quick < len(nums); quick++ {
		if nums[quick] == nums[slow-1] {
			continue
		}
		nums[slow] = nums[quick]
		slow++
	}
	return slow
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	idx := removeDuplicates(nums)
	fmt.Println(nums[:idx])
}
