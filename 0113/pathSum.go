func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	doPathSum(root, sum, []int{}, &res)
	return res
}

func doPathSum(root *TreeNode, sum int, path []int, res *[][]int) {
	if root == nil {
		return
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			path = append(path, root.Val)
			deepcopy := make([]int, len(path))
			copy(deepcopy, path)
			*res = append(*res, deepcopy)
		}
	}

	doPathSum(root.Left, sum, append(path, root.Val), res)
	doPathSum(root.Right, sum, append(path, root.Val), res)
}
