func nthUglyNumber(n int) int {
	factors := [3]int{2, 3, 5}
	// uglyNum = 2^count[0]*3^count[1]*5^count[2]
	count := []int{0, 0, 0}
	// candidates[0]: [1x2], 2x2, [2x2], 3x2, [3x2],[4x2], 5x2...
	// candidates[1]:  1x3, [1x3], 2x3,  2x3, [2x3], 3x3, [3x3]...
	// candidates[2]:  1x5,  1x5,  1x5, [1x5], 2x5,  2x5,  2x5...
	// 每个子列表都是一个丑数按顺序分别乘以 2，3，5
	candidates := [3]int{2, 3, 5}

	// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12...
	uglyNums := make([]int, n)
	uglyNums[0] = 1
	for i := 1; i < n; i++ {
		uglyNums[i] = min(candidates[0], candidates[1], candidates[2])
		for j := 0; j < 3; j++ {
			if uglyNums[i] == candidates[j] {
				count[j]++
				candidates[j] = uglyNums[count[j]] * factors[j]
			}
		}
	}
	return uglyNums[n-1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}
