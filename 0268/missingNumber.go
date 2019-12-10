func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		// 数字下标是从 [0, n-1]，数字是 [0, n]
		res = res ^ i ^ nums[i]
	}
	return res ^ len(nums)
}
