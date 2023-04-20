// 前置题 0206
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	preHead := &ListNode{Next: head}
	cur := preHead
	var preM, mPtr, nextN *ListNode
	for i := 1; cur != nil; i++ {
		if i == m {
			preM = cur
			mPtr = cur.Next
		}
		if i == n+1 {
			nextN = cur.Next
			cur.Next = nil
			break
		}
		cur = cur.Next
	}
	// pm   m           nn
	// 1 -> 2 -> 3 -> 4 5

	newHead := reverse(mPtr)
	// pm nh        m nn
	// 1  4 -> 3 -> 2 5
	// |            ^
	// -------------|

	preM.Next = newHead
	mPtr.Next = nextN

	return preHead.Next
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
