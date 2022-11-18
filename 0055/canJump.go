func canJump(nums []int) bool {
	canReach := 0
	for i := 0; i < len(nums) && i <= canReach; i++ {
		if i + nums[i] > canReach {
			canReach = i + nums[i]	
		}
		if canReach >= len(nums)-1 {
			return true
		}
	}
	return false
}
