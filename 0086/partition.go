func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	lowHead, highHead := &ListNode{}, &ListNode{}
	lowPtr, highPtr := lowHead, highHead
	curPtr := head
	for curPtr != nil {
		if curPtr.Val < x {
			lowPtr.Next = curPtr
			lowPtr = lowPtr.Next
		} else {
			highPtr.Next = curPtr
			highPtr = highPtr.Next
		}
		curPtr = curPtr.Next
	}
	lowPtr.Next = highHead.Next
	highPtr.Next = nil
	return lowHead.Next
}
