func generateTrees(n int) []*TreeNode {
	return generateBinarySearchTree(1, n)
}

func generateBinarySearchTree(start, end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		tree = append(tree, nil)
		return tree
	}

	for i := start; i <= end; i++ {
		left := generateBinarySearchTree(start, i-1)
		right := generateBinarySearchTree(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}
