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

// 类似题 0023
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := split(head)
	return merge2Lists(sortList(left), sortList(right))
}

func split(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	head2 := slow.Next
	slow.Next = nil
	return head, head2
}

func merge2Lists(left, right *ListNode) *ListNode {
	preHead := &ListNode{}
	cur := preHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}

	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return preHead.Next
}

func main() {
	head := &ListNode{
		5,
		&ListNode{
			7,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						8,
						nil,
					},
				},
			},
		},
	}
	newHead := sortList(head)
	printList(newHead)
}
