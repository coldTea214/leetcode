// 前置题 0094
// 递归版本
func preorderTraversal(root *TreeNode) []int {
	var res []int
	preorderTraversalHelper(root, &res)
	return res
}

func preorderTraversalHelper(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		preorderTraversalHelper(root.Left, res)
		preorderTraversalHelper(root.Right, res)
	}
}

// 迭代版本
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
	end := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return end
}

func preorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack Stack
	stack.push(root)

	for !stack.isEmpty() {
		node := stack.pop()
		if node == nil {
			continue
		}

		res = append(res, node.Val)
		stack.push(node.Right)
		stack.push(node.Left)
	}
	return res
}
