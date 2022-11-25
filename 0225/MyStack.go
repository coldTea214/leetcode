package main

import "fmt"

// 相关题 0232
type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Peek() int {
	return q.nums[0]
}

func (q *Queue) Empty() bool {
	return len(q.nums) == 0
}

func (q *Queue) Len() int {
	return len(q.nums)
}

//    1 2   3     pop 4     top
// a: 1 1,2 1,2,3     1,2,4 1,2,4
// b:             1,2
type MyStack struct {
	a, b *Queue
}

func Constructor() MyStack {
	return MyStack{a: NewQueue(), b: NewQueue()}
}

func (ms *MyStack) Push(x int) {
	if ms.a.Empty() {
		ms.a, ms.b = ms.b, ms.a
	}
	ms.a.Push(x)
}

func (ms *MyStack) Pop() int {
	if ms.a.Empty() {
		ms.a, ms.b = ms.b, ms.a
	}
	for ms.a.Len() != 1 {
		ms.b.Push(ms.a.Pop())
	}
	return ms.a.Pop()
}

func (ms *MyStack) Top() int {
	res := ms.Pop()
	ms.Push(res)
	return res
}

func (ms *MyStack) Empty() bool {
	return ms.a.Empty() && ms.b.Empty()
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	stack.Push(4)
	fmt.Println(stack.Top())
	// fmt.Println(stack.a.nums, stack.b.nums)
}
