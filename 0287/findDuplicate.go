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
