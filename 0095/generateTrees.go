func generateTrees(n int) []*TreeNode {
	return generateTreesHelper(1, n)
}

func generateTreesHelper(start, end int) []*TreeNode {
	trees := []*TreeNode{}
	if start > end {
		trees = append(trees, nil)
		return trees
	}

	for i := start; i <= end; i++ {
		left := generateTreesHelper(start, i-1)
		right := generateTreesHelper(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				trees = append(trees, root)
			}
		}
	}
	return trees
}
