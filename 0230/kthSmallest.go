func kthSmallest(root *TreeNode, k int) int {
	var res int
	inorder(root, &k, &res)
	return res
}

func inorder(node *TreeNode, k *int, res *int) {
	if node != nil && *k != 0 {
		inorder(node.Left, k, res)
		*k--
		if *k == 0 {
			*res = node.Val
			return
		}
		inorder(node.Right, k, res)
	}
}
