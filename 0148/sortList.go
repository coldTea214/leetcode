// 类似题 0023
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := split(head)
	return merge2Lists(sortList(left), sortList(right))
}

func split(head *ListNode) (left, right *ListNode) {
	slow, fast := head, head
	var preSlow *ListNode

	for fast != nil && fast.Next != nil {
		preSlow, slow = slow, slow.Next
		fast = fast.Next.Next
	}

	preSlow.Next = nil
	left, right = head, slow
	return
}

func merge2Lists(left, right *ListNode) *ListNode {
	cur := &ListNode{}
	preHead := cur
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}

	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return preHead.Next
}
