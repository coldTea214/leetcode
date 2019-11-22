/*
5: 0101
6: 0110
7: 0111
4: 0100 <-
*/
func rangeBitwiseAnd(m int, n int) int {
	var zeroBitNum uint
	for m >= 1 && n >= 1 {
		// 如果高位出现0，低位必然有0，当 m=n 时，已经过滤了所有含 0 的位
		if m == n {
			return m << zeroBitNum
		}

		m >>= 1
		n >>= 1
		zeroBitNum++
	}
	return 0
}
