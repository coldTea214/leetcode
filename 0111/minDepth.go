func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMin := minDepth(root.Left)
	rightMin := minDepth(root.Right)
	return min(leftMin, rightMin) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
