func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slowPtr, quickPtr := head, head

	for quickPtr != nil {
		if n < 0 {
			slowPtr = slowPtr.Next
		}

		n--
		quickPtr = quickPtr.Next
	}

	if n == 0 {
		return head.Next
	}
	slowPtr.Next = slowPtr.Next.Next
	return head
}
