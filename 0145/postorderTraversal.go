// 递归
func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorderTraversalHelper(root, &res)
	return res
}

func postorderTraversalHelper(root *TreeNode, res *[]int) {
	if root != nil {
		postorderTraversalHelper(root.Left, res)
		postorderTraversalHelper(root.Right, res)
		*res = append(*res, root.Val)
	}
}

// 迭代
type Stack struct {
	nodes []*TreeNode
}

func (s *Stack) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *Stack) top() *TreeNode {
	return s.nodes[len(s.nodes)-1]
}

func (s *Stack) push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) pop() {
	s.nodes = s.nodes[:len(s.nodes)-1]
}

func postorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack Stack
	stack.push(root)
	var pre, cur *TreeNode

	for !stack.isEmpty() {
		cur = stack.top()
		if cur == nil {
			stack.pop()
			continue
		}

		// 当前节点被读取的条件为：无左右孩子，或者上一次读取的为其左右孩子。
		if (cur.Left == nil && cur.Right == nil) ||
			(pre != nil && (pre == cur.Left || pre == cur.Right)) {
			res = append(res, cur.Val)
			stack.pop()
			pre = cur
		} else {
			stack.push(cur.Right)
			stack.push(cur.Left)
		}
	}

	return res
}
