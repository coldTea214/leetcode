func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	curPtr := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curPtr.Next = l1
			l1 = l1.Next
		} else {
			curPtr.Next = l2
			l2 = l2.Next
		}
		curPtr = curPtr.Next
	}

	if l1 != nil {
		curPtr.Next = l1
	}
	if l2 != nil {
		curPtr.Next = l2
	}
	return preHead.Next
}
