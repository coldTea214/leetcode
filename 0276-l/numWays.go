package main

import "fmt"

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	a, b, c := k, k*k, 0
	for i := 2; i < n; i++ {
		// 第n个栅栏如果和上一个不同颜色, 则有 b*(k-1) 种选择
		//              和上一个同颜色, 则有a*(k-1) 种
		c = b*(k-1) + a*(k-1)
		a, b = b, c
	}
	return b
}

func main() {
	fmt.Println(numWays(3, 2))
}
