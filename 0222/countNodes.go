package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 完全二叉树
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}

	low, high := 1, 1<<(level+1)
	for low < high {
		mid := (low + high) >> 1
		if mid <= 1<<level {
			low = mid + 1
		} else {
			bits := 1 << (level - 1)
			node := root
			for node != nil && bits > 0 {
				if bits&mid == 0 {
					node = node.Left
				} else {
					node = node.Right
				}
				bits >>= 1
			}
			if node == nil {
				high = mid
			} else {
				low = mid + 1
			}
		}
	}
	return low-1
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
