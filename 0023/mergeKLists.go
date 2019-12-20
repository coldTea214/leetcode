// 归并排序思想，复杂度 n*logk
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return doMerge(lists)
}

func doMerge(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return merge2Lists(lists[0], lists[1])
	}

	half := len(lists) >> 1
	return doMerge([]*ListNode{
		doMerge(lists[:half]),
		doMerge(lists[half:]),
	})
}

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	curPtr := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curPtr.Next = l1
			l1 = l1.Next
		} else {
			curPtr.Next = l2
			l2 = l2.Next
		}
		curPtr = curPtr.Next
	}

	if l1 != nil {
		curPtr.Next = l1
	}
	if l2 != nil {
		curPtr.Next = l2
	}
	return preHead.Next
}

