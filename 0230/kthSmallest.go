func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	inorder(root, k, &count, &res)
	return res
}

func inorder(node *TreeNode, k int, count *int, res *int) {
	if node != nil {
		inorder(node.Left, k, count, res)
		*count++
		if *count == k {
			*res = node.Val
			return
		}
		inorder(node.Right, k, count, res)
	}
}
