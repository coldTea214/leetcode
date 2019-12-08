func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 寻找中间结点
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	preMiddle := p1

	// 反转链表后半部分 1->2->3->4->5->6 to 1->2->3->6->5->4
	//                     p1    p2
	preCurrent := p1.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	p1 = head
	p2 = preMiddle.Next
	for p1 != preMiddle {
		if p1.Val == p2.Val {
			p1 = p1.Next
			p2 = p2.Next
		} else {
			return false
		}
	}
	if p1 == preMiddle {
		if p2 != nil && p1.Val != p2.Val {
			return false
		}
	}
	return true
}
