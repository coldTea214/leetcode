package main

import "fmt"

func verifyPreorder(preorder []int) bool {
	if len(preorder) < 2 {
		return true
	}
	var stack []int
	var min = -1 << 63
	stack = append(stack, preorder[0])
	for i := 1; i < len(preorder); i++ {
		// stack:
		// [5]
		// [5,2]
		// [5,2,1]
		// [5,3]
		for len(stack) > 0 && stack[len(stack)-1] < preorder[i] {
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		// 确保右子树大于根节点的值
		if preorder[i] < min {
			return false
		}

		stack = append(stack, preorder[i])
	}

	return true
}

func main() {
	fmt.Println(verifyPreorder([]int{5, 2, 1, 3, 6}))
}
