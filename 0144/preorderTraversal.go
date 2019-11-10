// 递归版本
func preorderTraversal(root *TreeNode) []int {
	var res []int
	doPreorderTraversal(root, &res)
	return res
}

func doPreorderTraversal(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		doPreorderTraversal(root.Left, res)
		doPreorderTraversal(root.Right, res)
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

func preorderTraversal(root *TreeNode) []int {
	var stack Stack
	stack.push(root)

	var res []int
	for !stack.isEmpty() {
		end := stack.pop()
		if end == nil {
			continue
		}

		res = append(res, end.Val)
		stack.push(end.Right)
		stack.push(end.Left)
	}
	return res
}
