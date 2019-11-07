import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxFromBeg := make([]int, len(prices)+1)
	for i := 1; i <= len(prices); i++ {
		if prices[i-1] < minPrice {
			minPrice = prices[i-1]
		}
		maxFromBeg[i] = max(prices[i-1]-minPrice, maxFromBeg[i-1])
	}

	maxPrice := 0
	maxFromEnd := make([]int, len(prices)+1)
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
		maxFromEnd[i] = max(maxFromEnd[i+1], maxPrice-prices[i])
	}

	res := 0
	for i := 0; i < len(prices); i++ {
		if maxFromBeg[i]+maxFromEnd[i] > res {
			res = maxFromBeg[i] + maxFromEnd[i]
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
