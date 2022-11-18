package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLevel := treeLevel(root.Left)
	rightLevel := treeLevel(root.Right)
	if leftLevel == rightLevel {
		// 左子树是完全二叉树
		return countNodes(root.Right) + 1<<leftLevel
	} else {
		// 右子树是完全二叉树
		return countNodes(root.Left) + 1<<rightLevel
	}
}

func treeLevel(root *TreeNode) int {
	level := 0
	for node := root; node != nil; node = node.Left {
		level++
	}
	return level
}

// 通用
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	var res int
	for len(queue) != 0 {
		node := queue[0]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
		res++
	}
	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	/*
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
	*/
	fmt.Println(countNodes(root))
}
