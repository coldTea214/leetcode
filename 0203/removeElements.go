func removeElements(head *ListNode, val int) *ListNode {
	preHead := ListNode{Next: head}

	cur := &preHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return preHead.Next
}
