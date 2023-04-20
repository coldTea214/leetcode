func pathSum(root *TreeNode, sum int) [][]int {
	var paths [][]int
	pathSumHelper(root, sum, []int{}, &paths)
	return paths
}

func pathSumHelper(root *TreeNode, sum int, path []int, paths *[][]int) {
	if root == nil {
		// append paths 相关逻辑不能放这, left/right nil 会入队2次
		return
	}

	sum -= root.Val
	path = append(path, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		deepcopy := make([]int, len(path))
		copy(deepcopy, path)
		*paths = append(*paths, deepcopy)
	}

	pathSumHelper(root.Left, sum, path, paths)
	pathSumHelper(root.Right, sum, path, paths)
}
