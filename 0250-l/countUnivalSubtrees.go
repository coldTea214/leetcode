package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const fakeVal = 1234567

func countUnivalSubtrees(root *TreeNode) int {
	var cnt int
	countUnivalSubtreesHelper(root, &cnt)
	return cnt
}

func countUnivalSubtreesHelper(root *TreeNode, cnt *int) (int, bool) {
	if root == nil {
		return fakeVal, true
	}
	lVal, lIsUnival := countUnivalSubtreesHelper(root.Left, cnt)
	rVal, rIsUnival := countUnivalSubtreesHelper(root.Right, cnt)
	isUnival := lIsUnival && rIsUnival &&
		(root.Val == lVal || lVal == fakeVal) && 
		(root.Val == rVal || rVal == fakeVal)
	if isUnival {
		*cnt++
	}
	return root.Val, isUnival
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(countUnivalSubtrees(root))
}
