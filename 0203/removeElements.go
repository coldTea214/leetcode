func removeElements(head *ListNode, val int) *ListNode {
	headPre := ListNode{Next: head}

	cur := &headPre
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return headPre.Next
}
