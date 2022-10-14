// 前置题 0137
func singleNumber(nums []int) []int {
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	// 获取最后一位不相同的 bit
	//  4: 0000 0100  5: 0000 0101
	// -4: 1111 1100 -5: 1111 1011
	diff &= -diff

	res := []int{0, 0}
	for _, num := range nums {
		if (num & diff) == 0 {
			res[0] ^= num
		} else { // the bit is set
			res[1] ^= num
		}
	}
	return res
}
