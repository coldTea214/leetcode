package main

import "fmt"

var inners = []string{"", "0", "1", "8"}
var pairs = [][]string{{"0", "0"}, {"1", "1"}, {"6", "9"}, {"8", "8"}, {"9", "6"}}

func strobogrammaticInRange(low string, high string) int {
	var ans int
	for _, inner := range inners {
		strobogrammaticInRangeHelper(low, high, inner, &ans)
	}
	return ans
}

func strobogrammaticInRangeHelper(low, high, num string, ans *int) {
	// 直接以字符串比较替代数字比较是不行的: "50" > "100"
	if len(low) <= len(num) && len(num) <= len(high) {
		if len(num) == len(low) && num < low ||
			len(num) == len(high) && num > high {
			return
		}
		if len(num) == 1 || num[0] != '0' {
			*ans++
		}
	}
	if len(num)+2 > len(high) {
		return
	}
	for _, pair := range pairs {
		strobogrammaticInRangeHelper(low, high, pair[0]+num+pair[1], ans)
	}
}

func main() {
	fmt.Println(strobogrammaticInRange("50", "100"))
}
