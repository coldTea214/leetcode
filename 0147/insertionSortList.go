func insertionSortList(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}

	cur := head
	for cur != nil && cur.Next != nil {
		needInsert := cur.Next
		if cur.Val <= needInsert.Val {
			cur = needInsert
			continue
		}

		cur.Next = needInsert.Next

		pre, next := headPre, headPre.Next
		for next.Val < needInsert.Val {
			pre = next
			next = next.Next
		}

		pre.Next = needInsert
		needInsert.Next = next
	}

	return headPre.Next
}
