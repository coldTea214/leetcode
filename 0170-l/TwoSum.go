package main

import "fmt"

type TwoSum struct {
	numCnt map[int]int
}

func Constructor() TwoSum {
	return TwoSum{
		numCnt: make(map[int]int),
	}
}

func (ts *TwoSum) Add(number int) {
	ts.numCnt[number]++
}

func (ts *TwoSum) Find(value int) bool {
	for n := range ts.numCnt {
		if value-n == n {
			if ts.numCnt[n] > 1 {
				return true
			}
		} else {
			if ts.numCnt[value-n] > 0 {
				return true
			}
		}
	}
	return false
}

func main() {
	ts := Constructor()
	ts.Add(3)
	ts.Add(1)
	ts.Add(2)
	fmt.Println(ts.Find(3))
}
