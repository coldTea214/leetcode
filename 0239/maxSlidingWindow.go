package main

import "fmt"

// 稀疏表
// leftMax:  i/k*k 到 i/k*k+i%k   区间求max
// rightMax: i     到 (i+k-1)/k*k 区间求max
// res:      i     到 i+k-1       区间求max，可以通过rightMax和leftMax错位取max求解
//
// leftMax:  0 max(0-1) max(0-2) 3        max(3-4) max(3-5) 6
// rightMax:            0        max(1-3) max(2-3) 3        max(4-6) max(5-6) 6
// res:                 max(0-2) max(1-3) max(2-4) max(3-5) max(4-6) max(5-6) 6
func maxSlidingWindow(nums []int, k int) []int {
	if k <= 1 {
		return nums
	}

	n := len(nums)
	leftMax := make([]int, n)
	for i := 0; i < n; i++ {
		if i%k == 0 {
			leftMax[i] = nums[i]
		} else {
			leftMax[i] = max(nums[i], leftMax[i-1])
		}
	}

	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if i%k == 0 {
			rightMax[i] = nums[i]
		} else {
			rightMax[i] = max(nums[i], rightMax[i+1])
		}
	}

	res := make([]int, n-k+1)
	for i := 0; i <= n-k; i++ {
		res[i] = max(rightMax[i], leftMax[i+k-1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
