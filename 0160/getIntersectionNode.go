func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	aLen, bLen := 1, 1
	tailA, tailB := headA, headB
	for tailA.Next != nil {
		tailA = tailA.Next
		aLen++
	}
	for tailB.Next != nil {
		tailB = tailB.Next
		bLen++
	}
	if tailA != tailB {
		return nil
	}

	if aLen < bLen {
		aLen, bLen = bLen, aLen
		headA, headB = headB, headA
	}
	for i := 1; i <= aLen-bLen; i++ {
		headA = headA.Next
	}
	for headA != headB {
		headA, headB = headA.Next, headB.Next
	}
	return headA
}
