// 本题核心是归纳场景，其实题目中都已经给出来了
// 1. 中序遍历只有一个错序：[1 3 2 4]，直接交换即可
// 2. 中序遍历能发现两个错序：[3 2 1]，则将前一个错序的较大者和后一个错序的较小者交换
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	inorderTraversal(root, &first, &second, &prev)
	first.Val, second.Val = second.Val, first.Val
}

func inorderTraversal(root *TreeNode, first, second, prev **TreeNode) {
	if root != nil {
		inorderTraversal(root.Left, first, second, prev)

		if *prev != nil && (*prev).Val > root.Val {
			if *first == nil {
				*first = *prev
			}
			// 当存在第二组错序的时候，second 的值，会被修改
			*second = root
		}
		*prev = root

		inorderTraversal(root.Right, first, second, prev)
	}
}
