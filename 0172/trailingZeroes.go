package main

import "fmt"

func trailingZeroes(n int) int {
	res := 0

	// 30: 7(6+1)
	// 含1个因子5: 6(5 10 15 20 25 30)
	// 含2个因子5: 1(25)
	for n >= 5 {
		n /= 5
		res += n
	}

	return res
}

func main() {
	fmt.Println(trailingZeroes(30))
}
