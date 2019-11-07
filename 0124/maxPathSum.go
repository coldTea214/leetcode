func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val
	findMaxSum(root, &maxSum)
	return maxSum
}

// 返回的是以 root 为"起点"，最大路径和
func findMaxSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	left := max(0, findMaxSum(root.Left, maxSum))
	right := max(0, findMaxSum(root.Right, maxSum))
	sum := left + root.Val + right
	if *maxSum < sum {
		*maxSum = sum
	}

	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}