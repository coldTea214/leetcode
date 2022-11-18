// 递归
func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorderTraversalHelper(root, &res)
	return res
}

func inorderTraversalHelper(root *TreeNode, res *[]int) {
	if root != nil {
		inorderTraversalHelper(root.Left, res)
		*res = append(*res, root.Val)
		inorderTraversalHelper(root.Right, res)
	}
}

// 迭代
type Stack struct {
	nodes []*TreeNode
}

func (s *Stack) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *Stack) push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) pop() *TreeNode {
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return node
}

func inorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack Stack

	for root != nil || !stack.isEmpty() {
		for root != nil {
			stack.push(root)
			root = root.Left
		}

		if !stack.isEmpty() {
			node := stack.pop()
			res = append(res, node.Val)
			root = node.Right
		}
	}
	return res
}
