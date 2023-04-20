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

// 迭代
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	preHead := &ListNode{0, head}
	cur := preHead
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return preHead.Next
}

// 递归
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 要么 head 重复了，那就删除 head
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	}

	// 要么 head 不重复，递归处理 head 后面的节点
	head.Next = deleteDuplicates(head.Next)
	return head
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
						4,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
		},
	}
	printList(deleteDuplicates(head))
}
