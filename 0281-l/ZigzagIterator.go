package main

import "fmt"

type ZigzagIterator struct {
	nums []int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	n1, n2 := len(v1), len(v2)
	n := n1
	if n1 < n2 {
		n = n2
	}
	nums := []int{}
	for i := 0; i < n; i++ {
		if i < n1 {
			nums = append(nums, v1[i])
		}
		if i < n2 {
			nums = append(nums, v2[i])
		}
	}
	return &ZigzagIterator{nums: nums}

}

func (zi *ZigzagIterator) next() int {
	if !zi.hasNext() {
		return 0
	}
	x := zi.nums[0]
	zi.nums = zi.nums[1:]
	return x
}

func (zi *ZigzagIterator) hasNext() bool {
	return len(zi.nums) > 0
}

func main() {
	v1 := []int{1, 2}
	v2 := []int{3, 4, 5, 6}
	zi := Constructor(v1, v2)
	for zi.hasNext() {
		fmt.Println(zi.next())
	}
}
