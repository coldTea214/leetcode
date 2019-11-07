type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Queue struct {
	nodes []*Node
}

func (q *Queue) isEmpty() bool {
	return len(q.nodes) == 0
}

func (q *Queue) length() int {
	return len(q.nodes)
}

func (q *Queue) enqueue(node *Node) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) dequeue() *Node {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func connect(root *Node) {
	var queue Queue
	queue.enqueue(root)
	for !queue.isEmpty() {
		// 记录 qLen 以确认每层的边界
		qLen := queue.length()
		var level []*Node
		for i := 1; i <= qLen; i++ {
			root = queue.dequeue()
			if root == nil {
				continue
			}
			level = append(level, root)
			queue.enqueue(root.Left)
			queue.enqueue(root.Right)
		}
		for i := 0; i < len(level)-1; i++ {
			level[i].Next = level[i+1]
		}
	}
}
