func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	fakeHead := &ListNode{Next: head}
	m++
	n++

	mPre, mPtr, nNext := split(fakeHead, m, n)
	// x->1 2->3->4 5->NULL

	h, e := reverse(mPtr)
	// x->1 4->3->2 5->NULL

	mPre.Next = h
	e.Next = nNext
	// x->1->4->3->2->5->NULL

	return fakeHead.Next
}

func split(head *ListNode, m, n int) (mPre, mPtr, nNext *ListNode) {
	for i := 1; head != nil; i++ {
		if i == m-1 {
			mPre = head
			mPtr = head.Next
		}
		if i == n {
			nNext = head.Next
			head.Next = nil
			break
		}
		head = head.Next
	}
	return
}

func reverse(head *ListNode) (h, e *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	var end *ListNode

	h, end = reverse(head.Next)
	end.Next = head
	e = head

	return
}
