func sumNumbers(root *TreeNode) int {
	res := 0
	sumNumbersHelper(root, 0, &res)
	return res
}

func sumNumbersHelper(root *TreeNode, num int, res *int) {
	if root == nil {
		return
	}

	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += num
		return
	}

	sumNumbersHelper(root.Left, num, res)
	sumNumbersHelper(root.Right, num, res)
}
