package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 寻找中间结点
	slow, quick := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}

	// 反转链表后半部分
	head2 := slow.Next
	var prev *ListNode
	for head2 != nil {
		next := head2.Next
		head2.Next = prev
		prev = head2
		head2 = next
	}
	// 1 -> 2 -> 3 2 <- 1
	// 1 -> 2 -> 3 3 <- 2 <- 1

	head2 = prev
	for head != nil && head2 != nil {
		if head.Val == head2.Val {
			head = head.Next
			head2 = head2.Next
		} else {
			return false
		}
	}
	return true
}

func main() {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					3,
					&ListNode{
						2,
						&ListNode{
							1,
							nil,
						},
					},
				},
			},
		},
	}
	fmt.Println(isPalindrome(head))
}
