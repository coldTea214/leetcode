package main

import "fmt"

func trailingZeroes(n int) int {
	res := 0

	// 150: 37(30 + 6 + 1)，分别为含 1、2、3 个因子 5 的个数
	for n >= 5 {
		n /= 5
		res += n
	}

	return res
}

func main() {
	fmt.Println(trailingZeroes(10))
}
