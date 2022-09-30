func insertionSortList(head *ListNode) *ListNode {
	preHead := &ListNode{Next: head}

	cur := head
	for cur != nil && cur.Next != nil {
		needInsert := cur.Next
		if cur.Val <= needInsert.Val {
			cur = needInsert
			continue
		}

		cur.Next = needInsert.Next

		pre, next := preHead, preHead.Next
		for next.Val < needInsert.Val {
			pre = next
			next = next.Next
		}

		pre.Next = needInsert
		needInsert.Next = next
	}

	return preHead.Next
}
