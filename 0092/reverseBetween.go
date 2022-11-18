func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	preHead := &ListNode{Next: head}

	preM, mPtr, nextN := split(preHead, m, n)
	// x->1 2->3->4 5->NULL

	newHead := reverse(mPtr)
	// x->1 4->3->2 5->NULL

	preM.Next = newHead
	mPtr.Next = nextN
	// x->1->4->3->2->5->NULL

	return preHead.Next
}

func split(head *ListNode, m, n int) (preM, mPtr, nextN *ListNode) {
	for i := 1; head != nil; i++ {
		if i == m {
			preM = head
			mPtr = head.Next
		}
		if i == n+1 {
			nextN = head.Next
			head.Next = nil
			break
		}
		head = head.Next
	}
	return
}

func reverse(head *ListNode) (newHead *ListNode) {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
