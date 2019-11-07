func maxDepth(root *TreeNode) int {
	var max int
	inorderTraversal(root, 0, &max)
	return max
}

func inorderTraversal(root *TreeNode, depth int, max *int) {
	if root == nil {
		if depth > *max {
			*max = depth
		}
		return
	}

	inorderTraversal(root.Left, depth+1, max)
	inorderTraversal(root.Right, depth+1, max)
}
