func canJump(nums []int) bool {
	canJumpTo := make([]bool, len(nums))
	canJumpTo[0] = true
	for i := 0; i < len(nums) && canJumpTo[i]; i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			canJumpTo[i+j] = true
		}
	}
	return canJumpTo[len(nums)-1]
}
