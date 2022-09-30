// 递归版本
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}

func swapPairs2(head *ListNode) *ListNode {
	preHead := &ListNode{0, head}
    tmp := preHead
    for tmp.Next != nil && tmp.Next.Next != nil {
        node1 := tmp.Next
        node2 := tmp.Next.Next
        tmp.Next = node2
        node1.Next = node2.Next
        node2.Next = node1
        tmp = node1
    }
    return preHead.Next
}

