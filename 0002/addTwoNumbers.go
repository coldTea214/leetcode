// 执行用时 : 16 ms , 在所有 golang 提交中击败了 62.61% 的用户
// 内存消耗 : 4.9 MB , 在所有 golang 提交中击败了 94.01% 的用户

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{}
	curPtr := fakeHead

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

	return fakeHead.Next
}
