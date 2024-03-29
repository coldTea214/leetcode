package main

import (
	"fmt"
)

// 前置题 0055
func jump(nums []int) int {
	lastReach := 0
	canReach := 0
	steps := 0
	for i := 0; i < len(nums)-1 && i <= canReach; i++ {
		if canReach < i+nums[i] {
			canReach = i + nums[i]
		}
		if i == lastReach {
			lastReach = canReach
			steps++
		}
	}
	return steps
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
