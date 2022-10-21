package main

import (
	"fmt"
	"strings"
)

var pairs = [][2]byte{{'0', '0'}, {'1', '1'}, {'6', '9'}, {'8', '8'}, {'9', '6'}}
var base1 = []string{"0", "1", "8"}
var base2 = []string{"00", "11", "69", "88", "96"}

func findStrobogrammatic(n int) []string {
	if n == 2 {
		return base2[1:]
	}
	return findStrobogrammaticHelper(n, n)
}

func findStrobogrammaticHelper(n, k int) []string {
	if k == 1 {
		return base1
	}
	if k == 2 {
		return base2
	}

	innerNums := findStrobogrammaticHelper(n, k-2)
	nums := []string{}
	for _, innerNum := range innerNums {
		for _, pair := range pairs {
			if k == n && pair[0] == '0' {
				continue
			}
			num := &strings.Builder{}
			num.WriteByte(pair[0])
			num.WriteString(innerNum)
			num.WriteByte(pair[1])
			nums = append(nums, num.String())
		}
	}
	return nums
}

func main() {
	fmt.Println(findStrobogrammatic(1))
	fmt.Println(findStrobogrammatic(2))
	fmt.Println(findStrobogrammatic(3))
	fmt.Println(findStrobogrammatic(4))
}
