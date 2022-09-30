func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}
	return reverseHelper(head, k)
}

func reverseHelper(head *ListNode, k int) *ListNode {
	tail, needReverse := getTail(head, k)
	if needReverse {
		tailNext := tail.Next
		// k = 3
		// 3 -> 2 -> 1 -> 4 -> 5 -> 6 -> 7
		//                h         t    tn
		head, tail = reverseList(head, tail)
		// 3 -> 2 -> 1 -> 6 -> 5 -> 4    7
		//                h         t    tn
		tail.Next = reverseHelper(tailNext, k)
	}

	return head
}

func getTail(head *ListNode, k int) (*ListNode, bool) {
	for i := 1; i < k; i++ {
		if head != nil {
			head = head.Next
		} else {
			return nil, false
		}
	}
	return head, head != nil
}

// input:  1 -> 2 -> 3 ->
// output: 3 -> 2 -> 1
func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	pre, cur := head, head.Next
	end := tail.Next
	for cur != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
