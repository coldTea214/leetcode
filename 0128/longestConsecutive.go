func longestConsecutive(nums []int) int {
	var res int
	// lenWithEndpoint 记录以 num 为端点, 连续串的长度
	lenWithEndpoint := make(map[int]int)
	for _, num := range nums {
		if lenWithEndpoint[num] != 0 {
			continue
		}

		var left, right int
		if lenWithEndpoint[num-1] > 0 {
			left = lenWithEndpoint[num-1]
		}
		if lenWithEndpoint[num+1] > 0 {
			right = lenWithEndpoint[num+1]
		}

		sum := left + right + 1
		if sum > res {
			res = sum
		}

		// 这里放 -1 也无所谓:
		// 如果 num 当前没跟任何点连起来, 它的值是什么不重要
		// 如果 num 不是端点, 它的值是什么也不重要, <=0 即可
		lenWithEndpoint[num] = -1
		lenWithEndpoint[num-left] = sum
		lenWithEndpoint[num+right] = sum
	}
	return res
}
