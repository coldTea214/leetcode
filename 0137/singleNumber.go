// 后置题 0260
// 针对 1 个数字 n：
// 出现 1 次时，ones = n, twos = 0
// 出现 2 次时，ones = 0, twos = n
// 出现 3 次时，ones = 0, twos = 0
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

// 上面的算式需要推演，更直白的方法：
// 1. hash 统计计数，时间复杂度O(n)，空间复杂度O(n)
// 2. 统计数字0-31位，每1位相加模3再合入结果，时间复杂度O(n)，空间复杂度O(1)
// 方法2形如:
func singleNumber2(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
