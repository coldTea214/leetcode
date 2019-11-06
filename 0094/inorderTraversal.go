// 递归版本
func inorderTraversal(root *TreeNode) []int {
	var res []int
	doInorderTraversal(root, &res)
	return res
}

func doInorderTraversal(root *TreeNode, res *[]int) {
	if root != nil {
		doInorderTraversal(root.Left, res)
		*res = append(*res, root.Val)
		doInorderTraversal(root.Right, res)
	}
}

// 非递归版本
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

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack Stack
	for root != nil || !stack.isEmpty() {
		for root != nil {
			stack.push(root)
			root = root.Left
		}

		if !stack.isEmpty() {
			top := stack.pop()
			res = append(res, top.Val)
			root = top.Right
		}
	}
	return res
}
