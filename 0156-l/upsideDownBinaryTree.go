func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}

	leftTree, rightTree := root.Left, root.Right
	root.Left, root.Right = nil, nil
	leftTreeRoot := upsideDownBinaryTree(leftTree)
	leftTree.Right, leftTree.Left = root, rightTree
	return leftTreeRoot
}
