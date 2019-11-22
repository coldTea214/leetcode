type Queue struct {
	nodes []*TreeNode
}

func (q *Queue) isEmpty() bool {
	return len(q.nodes) == 0
}

func (q *Queue) length() int {
	return len(q.nodes)
}

func (q *Queue) enqueue(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) dequeue() *TreeNode {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func rightSideView(root *TreeNode) []int {
	var queue Queue
	var res []int
	queue.enqueue(root)
	for !queue.isEmpty() {
		// 记录 qLen 以确认每层的边界
		qLen := queue.length()
		var rightVal int
		for i := 1; i <= qLen; i++ {
			root = queue.dequeue()
			if root == nil {
				continue
			}
			rightVal = root.Val
			queue.enqueue(root.Left)
			queue.enqueue(root.Right)
		}
		if rightVal != 0 {
			res = append(res, rightVal)
		}
	}
	return res
}
