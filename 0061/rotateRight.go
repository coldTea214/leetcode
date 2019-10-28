func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	tairPtr := head
	count := 1
	for tairPtr.Next != nil {
		tairPtr = tairPtr.Next
		count++
	}
	k %= count
	if k == 0 {
		return head
	}

	curPtr := head
	for i := 1; i < count-k; i++ {
		curPtr = curPtr.Next
	}
	newHead := curPtr.Next
	curPtr.Next = nil
	tairPtr.Next = head

	return newHead
}
