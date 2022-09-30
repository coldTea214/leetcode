func flatten(root *TreeNode) {
	flattenHelper(root)
}

// 返回的是 flatten 后的最后一个结点
func flattenHelper(root *TreeNode) *TreeNode {
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return flattenHelper(root.Right)
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return flattenHelper(root.Right)
	}

	lastNode := flattenHelper(root.Right)
	flattenHelper(root.Left).Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return lastNode
}
