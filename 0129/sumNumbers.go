func sumNumbers(root *TreeNode) int {
	res := 0
	calculateSum(root, 0, &res)
	return res
}

func calculateSum(root *TreeNode, num int, res *int) {
	if root == nil {
		return
	}

	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += num
		return
	}

	calculateSum(root.Left, num, res)
	calculateSum(root.Right, num, res)
}
