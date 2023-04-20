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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}
	return reverseHelper(head, k)
}

func reverseHelper(head *ListNode, k int) *ListNode {
	tail := head
	for i := 1; i < k && tail != nil; i++ {
		tail = tail.Next
	}
	// k = 3
	// h         t
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7

	if tail != nil {
		tailNext := tail.Next
		tail.Next = nil
		// h         t  tn
		// 1 -> 2 -> 3  4 -> 5 -> 6 -> 7

		head, tail = reverseList(head)
		// h         t tn
		// 3 -> 2 -> 1 4 -> 5 -> 6 -> 7

		tail.Next = reverseHelper(tailNext, k)
	}

	return head
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev, head
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
	printList(reverseKGroup(head, 2))
}
