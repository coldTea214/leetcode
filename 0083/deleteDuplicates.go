func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, quick := head, head.Next
	for ; quick != nil; quick = quick.Next {
		if quick.Val != slow.Val {
			slow.Next = quick
			slow = quick
		}
	}
	slow.Next = nil
	return head
}
