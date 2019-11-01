func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, quick := head, head.Next
	for quick != nil {
		if quick.Val != slow.Val {
			slow.Next = quick
			slow = slow.Next
		}
		quick = quick.Next
	}
	slow.Next = nil
	return head
}
