package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const invalidVal = -1 << 31

func longestConsecutive(root *TreeNode) int {
	var maxLen int
	longestConsecutiveHelper(root, &maxLen)
	return maxLen
}

func longestConsecutiveHelper(root *TreeNode, maxLen *int) (int, int) {
	if root == nil {
		return 0, invalidVal
	}

	lLen, lVal := longestConsecutiveHelper(root.Left, maxLen)
	rLen, rVal := longestConsecutiveHelper(root.Right, maxLen)

	length := 1
	if root.Val == lVal-1 {
		length = lLen + 1
	}
	if root.Val == rVal-1 {
		if rLen+1 > length {
			length = rLen + 1
		}
	}
	if length > *maxLen {
		*maxLen = length
	}
	return length, root.Val
}

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}
	fmt.Println(longestConsecutive(root))
}
