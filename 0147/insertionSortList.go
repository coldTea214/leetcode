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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	preHead := &ListNode{Next: head}
	// cur 之前的链表已有序
	cur := head
	for cur.Next != nil {
		// loop1: 5 7 3 4 8
		//        5 7 3 4 8
		//        3 5 7 4 8
		//        3 4 5 7 8
		needInsert := cur.Next
		if cur.Val <= needInsert.Val {
			cur = needInsert
			continue
		}
		// loop2: 
		//      c    n
		// 5 -> 7 -> 3 -> 4 -> 8

		cur.Next = needInsert.Next

		preInsert, nextInsert := preHead, preHead.Next
		for nextInsert.Val < needInsert.Val {
			preInsert = nextInsert
			nextInsert = nextInsert.Next
		}
		preInsert.Next = needInsert
		needInsert.Next = nextInsert
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
	insertionSortList(head)
}
