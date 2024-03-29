// 其实画个图更直白，有空再画
// 假设环的入口为点 a，快慢指针第一次相遇时的点为 b
// 设起点到 a 距离为 x，a 到 b 为 y，b 到 a 为 z
// 则 2 * (x + y) = x + n(y + z) + y, 得 x = z + (n-1)(y+z)
// 相遇后，两个慢指针分别从起点、b 出发，则下次相遇点为 a
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			break
		}
	}
	if slow != quick {
		return nil
	}

	cur := head
	for slow != cur {
		slow = slow.Next
		cur = cur.Next
	}
	return slow
}
