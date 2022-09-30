func canJump(nums []int) bool {
	canJumpTo := 0
	for i := 0; i < len(nums) && i <= canJumpTo; i++ {
		if i + nums[i] > canJumpTo {
			canJumpTo = i + nums[i]	
		}
		if canJumpTo >= len(nums)-1 {
			return true
		}
	}
	return false
}
