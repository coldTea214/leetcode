func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findAncestor(root, p.Val, q.Val)
}

func findAncestor(root *TreeNode, p, q int) *TreeNode {
	r := root.Val
	if p < r && q < r {
		return findAncestor(root.Left, p, q)
	} else if r < p && r < q {
		return findAncestor(root.Right, p, q)
	}
	return root
}
