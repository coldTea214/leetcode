// 前置题 0041
// 后置题 0287
func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		// 数字下标是从 [0, n-1]，数字是 [0, n]
		// 另, 此解法当且"仅当缺失1个数"时可用
		res = res ^ i ^ nums[i]
	}
	return res ^ len(nums)
}
