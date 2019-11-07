func sortedListToBST(head *ListNode) *TreeNode {
	return doSortedListToBST(head, nil)
}

func doSortedListToBST(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}

	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow

	return &TreeNode{
		Val:   mid.Val,
		Left:  doSortedListToBST(begin, mid),
		Right: doSortedListToBST(mid.Next, end),
	}
}
