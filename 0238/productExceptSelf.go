func productExceptSelf(nums []int) []int {
	// nums: [2, 3, 4, 5]
	res := make([]int, len(nums))

	left := 1
	for i := 0; i < len(nums); i++ {
		res[i] = left
		left *= nums[i]
	}
	// res: [1, 2, 2*3, 2*3*4]

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	// res: [1*(3*4*5), 2*(4*5), 2*3*(5), 2*3*4]

	return res
}
