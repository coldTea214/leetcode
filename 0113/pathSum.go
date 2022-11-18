func pathSum(root *TreeNode, sum int) [][]int {
	var paths [][]int
	pathSumHelper(root, sum, []int{}, &paths)
	return paths
}

func pathSumHelper(root *TreeNode, sum int, path []int, paths *[][]int) {
	if root == nil {
		return
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			path = append(path, root.Val)
			deepcopy := make([]int, len(path))
			copy(deepcopy, path)
			*paths = append(*paths, deepcopy)
		}
	}

	pathSumHelper(root.Left, sum, append(path, root.Val), paths)
	pathSumHelper(root.Right, sum, append(path, root.Val), paths)
}
