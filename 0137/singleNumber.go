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
