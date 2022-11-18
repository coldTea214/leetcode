func grayCode(n int) []int {
	return grayCodeHelper(n, 1, []int{0})
}

func grayCodeHelper(n, base int, nums []int) []int {
	if n == 0 {
		return nums
	}

	tmp := make([]int, len(nums))
	for i := range nums {
		/*
			每一轮 +base 本身都不会产生进位
			再通过逆序确保 +base 不会在两轮间(2->6)产生进位
			nums:
			[0]
			[0 1]
			[0 1 3 2]
			[0 1 3 2 6 7 5 4]
		*/
		tmp[len(nums)-i-1] = nums[i] + base
	}

	return grayCodeHelper(n-1, base*2, append(nums, tmp...))
}
