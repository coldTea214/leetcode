func hasCycle(head *ListNode) bool {
	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			return true
		}
	}
	return false
}
