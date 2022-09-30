# 解法一：只适用于完美二叉树
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	for levelFirstNode := root; levelFirstNode.Left != nil; levelFirstNode = levelFirstNode.Left {
		for node := levelFirstNode; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

# 解法二：层次遍历的变种
