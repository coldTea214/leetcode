func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	var res int
	for len(queue) != 0 {
		node := queue[0]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
		res++
	}
	return res
}
