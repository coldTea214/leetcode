func sortedListToBST(head *ListNode) *TreeNode {
	return sortedListToBSTHelper(head, nil)
}

func sortedListToBSTHelper(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	quick, slow := begin, begin
	for quick != end && quick.Next != end {
		quick = quick.Next.Next
		slow = slow.Next
	}

	mid := slow
	return &TreeNode{
		Val:   mid.Val,
		Left:  sortedListToBSTHelper(begin, mid),
		Right: sortedListToBSTHelper(mid.Next, end),
	}
}
