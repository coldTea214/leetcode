// 可以直接用复制图那题的解法
// 不过也有更简单的解法
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	copyHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		copyNode := cur.Next
		cur.Next = cur.Next.Next
		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}
	}
	return copyHead
}
