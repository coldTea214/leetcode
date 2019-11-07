func flatten(root *TreeNode) {
	doFlatten(root)
}

// 返回的是 flatten 后的最后一个结点
func doFlatten(root *TreeNode) *TreeNode {
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return doFlatten(root.Right)
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return doFlatten(root.Right)
	}

	res := doFlatten(root.Right)
	doFlatten(root.Left).Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return res
}
