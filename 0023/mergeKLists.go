// 这道题的直观解法应该是用一个大小为k的小根堆，每次取队列首结点处理，复杂度是n*logk，n为总结点个数
// 但 golang 并没有对应算法库，退而求其次用归并排序来做，复杂度也是n*logk

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return doMerge(lists)
}

func doMerge(lists []*ListNode) *ListNode {
	listNum := len(lists)

	if listNum == 1 {
		return lists[0]
	}

	if listNum == 2 {
		return merge2Lists(lists[0], lists[1])
	}

	half := listNum >> 1
	return doMerge([]*ListNode{
		doMerge(lists[:half]),
		doMerge(lists[half:]),
	})
}

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{}
	curPtr := fakeHead
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
	return fakeHead.Next
}
