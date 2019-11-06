func isValidBST(root *TreeNode) bool {
	res := inorderTraversal(root)
	for i := 1; i < len(res); i++ {
		if res[i-1] >= res[i] {
			return false
		}
	}
	return true
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	doInorderTraversal(root, &res)
	return res
}

func doInorderTraversal(root *TreeNode, res *[]int) {
	if root != nil {
		doInorderTraversal(root.Left, res)
		*res = append(*res, root.Val)
		doInorderTraversal(root.Right, res)
	}
}
