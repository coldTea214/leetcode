func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 获取 list 的长度 size
	cur := head
	size := 0
	for cur != nil {
		cur = cur.Next
		size++
	}

	// size 为奇数， cur 指向 list 的中间节点
	// size 为偶数， cur 指向 list 前一半的最后一个节点
	cur = head
	for i := 0; i < (size-1)/2; i++ {
		cur = cur.Next
	}
	// head -> 1 -> 2 -> 3 -> 4 -> 5 -> 6
	//                   ^
	//                   |
	//                  cur

	next := cur.Next
	for next != nil {
		temp := next.Next
		next.Next = cur
		cur = next
		next = temp
	}
	end := cur
	// head -> 1 -> 2 -> 3 <-> 4 <- 5 <- 6 <- end

	for head != end {
		hNext := head.Next
		eNext := end.Next
		head.Next = end
		end.Next = hNext
		head = hNext
		end = eNext
	}

	end.Next = nil
}
