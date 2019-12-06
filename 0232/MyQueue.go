type Stack struct {
	data []int
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x
}

func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

type MyQueue struct {
	a, b Stack
}

func Constructor() MyQueue {
	return MyQueue{
		a: Stack{},
		b: Stack{},
	}
}

func (q *MyQueue) Push(x int) {
	q.a.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.b.Empty() {
		for !q.a.Empty() {
			q.b.Push(q.a.Pop())
		}
	}
	return q.b.Pop()
}

func (q *MyQueue) Peek() int {
	if q.b.Empty() {
		for !q.a.Empty() {
			q.b.Push(q.a.Pop())
		}
	}
	return q.b.Top()
}

func (q *MyQueue) Empty() bool {
	return q.a.Empty() && q.b.Empty()
}
