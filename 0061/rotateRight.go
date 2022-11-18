func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	tair := head
	count := 1
	for tair.Next != nil {
		tair = tair.Next
		count++
	}
	k %= count
	if k == 0 {
		return head
	}

	cur := head
	for i := 1; i < count-k; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	tair.Next = head

	return newHead
}
