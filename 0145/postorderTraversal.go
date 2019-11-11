// 递归版本
func postorderTraversal(root *TreeNode) []int {
	var res []int
	doPostorderTraversal(root, &res)
	return res
}

func doPostorderTraversal(root *TreeNode, res *[]int) {
	if root != nil {
		doPostorderTraversal(root.Left, res)
		doPostorderTraversal(root.Right, res)
		*res = append(*res, root.Val)
	}
}

// 迭代版本
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

// 当前节点被读取的条件为：无左右孩子，或者上一次读取的为其左右孩子。
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack Stack
	stack.push(root)
	var pre, cur *TreeNode
	for !stack.isEmpty() {
		cur = stack.top()
		if (cur.Left == nil && cur.Right == nil) ||
			(pre != nil && (pre == cur.Left || pre == cur.Right)) {
			res = append(res, cur.Val)
			stack.pop()
			pre = cur
		} else {
			if cur.Right != nil {
				stack.push(cur.Right)
			}
			if cur.Left != nil {
				stack.push(cur.Left)
			}
		}
	}

	return res
}
