func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	if k >= len(prices) {
		return profits(prices)
	}

	// local[i][j]， 到达第 i 天时最多可进行 j 次交易 且最后一次交易在最后一天卖出 的最大利润
	// global[i][j]，到达第 i 天时最多可进行 j 次交易 的最大利润
	// local[i][j]  = max(global[i-1][j-1] + max(diff, 0), local[i-1][j] + diff)
	// global[i][j] = max(local[i][j], global[i-1][j])
	local := make([]int, k+1)
	global := make([]int, k+1)
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			local[j] = max(global[j-1]+max(diff, 0), local[j]+diff)
			global[j] = max(local[j], global[j])
		}
	}
	return global[k]
}

func profits(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
