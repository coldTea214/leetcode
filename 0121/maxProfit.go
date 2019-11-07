import "math"

func maxProfit(prices []int) int {
	max, min := 0, math.MaxInt32
	for _, price := range prices {
		if price < min {
			min = price
		} else {
			if price-min > max {
				max = price - min
			}
		}
	}
	return max
}
