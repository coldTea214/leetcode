import "math"

func maxProfit(prices []int) int {
	profit, min := 0, math.MaxInt32
	for _, price := range prices {
		if price < min {
			min = price
		} else {
			if price-min > profit {
				profit = price - min
			}
		}
	}
	return profit
}
