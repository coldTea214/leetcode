type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	res := BSTIterator{
		stack: []*TreeNode{},
	}
	res.pushUntilMin(root)
	return res
}

func (iter *BSTIterator) Next() int {
	size := len(iter.stack)
	var top *TreeNode
	iter.stack, top = iter.stack[:size-1], iter.stack[size-1]
	iter.pushUntilMin(top.Right)
	return top.Val
}

func (iter *BSTIterator) HasNext() bool {
	return len(iter.stack) > 0
}

func (iter *BSTIterator) pushUntilMin(root *TreeNode) {
	for root != nil {
		iter.stack = append(iter.stack, root)
		root = root.Left
	}
}
