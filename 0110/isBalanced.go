func isBalanced(root *TreeNode) bool {
	_, isBalanced := doIsBalanced(root)
	return isBalanced
}

func doIsBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := doIsBalanced(root.Left)
	rightDepth, rightIsBalanced := doIsBalanced(root.Right)

	if leftIsBalanced && rightIsBalanced &&
		-1 <= leftDepth-rightDepth && leftDepth-rightDepth <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
