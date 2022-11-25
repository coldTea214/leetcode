package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%v ", cur.Val)
	}
	fmt.Println()
}

func findMiddle(head *ListNode) {
	if head == nil {
		return
	}
	slow, quick := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	//           s
	// 1 -> 2 -> 3 -> 4 -> 5
	// h         s 
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6
	// tips: 如果上述判断条件是 for quick != nil && quick.Next != nil, 则 slow 位置略有不同
	//           s
	// 1 -> 2 -> 3 -> 4 -> 5
	//                s
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func main() {
	_ = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
}
