func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	curPtr := preHead

	var carry int
	for l1 != nil || l2 != nil {
		var digit1, digit2 int
		if l1 != nil {
			digit1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			digit2 = l2.Val
			l2 = l2.Next
		}

		sum := (digit1 + digit2 + carry) % 10
		carry = (digit1 + digit2 + carry) / 10

		curPtr.Next = &ListNode{
			Val: sum,
		}
		curPtr = curPtr.Next
	}
	if carry != 0 {
		curPtr.Next = &ListNode{
			Val: carry,
		}
	}

	return preHead.Next
}
