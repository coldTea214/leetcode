package main

import "fmt"

// 相关题 0225
type Stack struct {
	nums []int
}

func (s *Stack) Push(x int) {
	s.nums = append(s.nums, x)
}

func (s *Stack) Pop() int {
	x := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return x
}

func (s *Stack) Top() int {
	return s.nums[len(s.nums)-1]
}

func (s *Stack) Empty() bool {
	return len(s.nums) == 0
}

//    1 2   3     pop 4   peek
// a: 1 1,2 1,2,3     4   4
// b:             3,2 3,2 3,2
type MyQueue struct {
	a, b Stack
}

func Constructor() MyQueue {
	return MyQueue{a: Stack{}, b: Stack{}}
}

func (mq *MyQueue) Push(x int) {
	mq.a.Push(x)
}

func (mq *MyQueue) Pop() int {
	if mq.b.Empty() {
		for !mq.a.Empty() {
			mq.b.Push(mq.a.Pop())
		}
	}
	return mq.b.Pop()
}

func (mq *MyQueue) Peek() int {
	if mq.b.Empty() {
		for !mq.a.Empty() {
			mq.b.Push(mq.a.Pop())
		}
	}
	return mq.b.Top()
}

func (mq *MyQueue) Empty() bool {
	return mq.a.Empty() && mq.b.Empty()
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	queue.Push(4)
	fmt.Println(queue.Peek())
	// fmt.Println(queue.a.nums, queue.b.nums)
}
