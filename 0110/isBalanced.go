func isBalanced(root *TreeNode) bool {
	_, isBalanced := isBalancedHelper(root)
	return isBalanced
}

func isBalancedHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := isBalancedHelper(root.Left)
	rightDepth, rightIsBalanced := isBalancedHelper(root.Right)

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
