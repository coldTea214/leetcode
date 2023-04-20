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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{Next: head}
	slow, quick := preHead, preHead
	for i := 1; i <= n+1 && quick != nil; i++ {
		quick = quick.Next
	}
	for quick != nil {
		slow, quick = slow.Next, quick.Next
	}

	slow.Next = slow.Next.Next
	return preHead.Next
}

func main() {
	head := &ListNode{
		1,
		nil,
	}
	printList(removeNthFromEnd(head, 1))

	head = &ListNode{
		1,
		&ListNode{
			2,
			nil,
		},
	}
	printList(removeNthFromEnd(head, 1))

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
						nil,
					},
				},
			},
		},
	}
	printList(removeNthFromEnd(head, 2))
}
