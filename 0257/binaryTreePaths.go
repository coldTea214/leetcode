import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	res := []string{}
	left := binaryTreePaths(root.Left)
	for i := 0; i < len(left); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+left[i])
	}
	right := binaryTreePaths(root.Right)
	for i := 0; i < len(right); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+right[i])
	}
	return res
}
