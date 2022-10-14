package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	// 缓存已经计算过的字符串
	calculated := make(map[string][]int)
	return diffWaysToComputeHelper(input, calculated)
}

func diffWaysToComputeHelper(s string, calculated map[string][]int) []int {
	if res, ok := calculated[s]; ok {
		return res
	}

	var res []int
	for i := range s {
		if s[i] == '+' || s[i] == '-' || s[i] == '*' {
			// 把 s[i] 作为最后一个运算的运算符
			for _, left := range diffWaysToComputeHelper(s[:i], calculated) {
				for _, right := range diffWaysToComputeHelper(s[i+1:], calculated) {
					res = append(res, operate(left, right, s[i]))
				}
			}
		}
	}

	// s 中不存在运算符
	if len(res) == 0 {
		tmp, _ := strconv.Atoi(s)
		res = append(res, tmp)
	}

	calculated[s] = res
	return res
}

func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return a * b
	}
}

func main() {
	// [-34 -10 -14 -10 10]
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
