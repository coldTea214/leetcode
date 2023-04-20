package main

import "fmt"

// 前置题 0268, 这题就没有"仅缺失1个数"条件, 输入为 [2,2,2,2,2] 就没法用 ^ 处理
// 用 0041 的解法来没问题, 但题目中又有"不修改数组"的限制
// 将这些数字想象成链表中的结点，数组中数字代表下一个结点的数组下标，找重复的数字就是找链表中成环的那个点
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
