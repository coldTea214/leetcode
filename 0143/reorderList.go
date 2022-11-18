package main

import "fmt"

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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, quick := head, head
	// tips: 
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	if slow == nil || slow == quick {
		return
	}

	head2 := slow.Next
	slow.Next = nil
	//           s
	// 1 -> 2 -> 3  4 -> 5
	// h         s 
	// 1 -> 2 -> 3  4 -> 5 -> 6
	// tips: 如果上述判断条件是 for quick != nil && quick.Next != nil, 则 slow 位置略有不同
	//           s
	// 1 -> 2 -> 3  4 -> 5
	//                s
	// 1 -> 2 -> 3 -> 4  5 -> 6

	var prev *ListNode
	for head2 != nil {
		next := head2.Next
		head2.Next = prev
		prev = head2
		head2 = next
	}
	head2 = prev
	// h                      h2
	// 1 -> 2 -> 3  4 <- 5 <- 6

	for head != nil && head2 != nil {
		hNext := head.Next
		eNext := head2.Next
		head.Next = head2
		head2.Next = hNext
		head = hNext
		head2 = eNext
	}
}

func main() {
	head := &ListNode{
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
	reorderList(head)
	printList(head)

	head = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						&ListNode{
							6,
							nil,
						},
					},
				},
			},
		},
	}
	reorderList(head)
	printList(head)
}
