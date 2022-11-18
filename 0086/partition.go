func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	lowHead, highHead := &ListNode{}, &ListNode{}
	low, high := lowHead, highHead
	cur := head
	for cur != nil {
		if cur.Val < x {
			low.Next = cur
			low = low.Next
		} else {
			high.Next = cur
			high = high.Next
		}
		cur = cur.Next
	}
	low.Next = highHead.Next
	high.Next = nil
	return lowHead.Next
}
