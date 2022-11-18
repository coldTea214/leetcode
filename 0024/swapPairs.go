// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}

// 迭代
func swapPairs2(head *ListNode) *ListNode {
	preHead := &ListNode{0, head}
    cur := preHead
    for cur.Next != nil && cur.Next.Next != nil {
        node1 := cur.Next
        node2 := cur.Next.Next
        cur.Next = node2
        node1.Next = node2.Next
        node2.Next = node1
        cur = node1
    }
    return preHead.Next
}

