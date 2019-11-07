func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return doIsSymmetric(root.Left, root.Right)
}
func doIsSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	if p.Val != q.Val {
		return false
	}
	return doIsSymmetric(p.Left, q.Right) && doIsSymmetric(p.Right, q.Left)
}
