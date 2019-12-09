func singleNumber(nums []int) []int {
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	// 获取最后一位不相同的 bit
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
