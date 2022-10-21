package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	begin, end := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		begin[i] = intervals[i][0]
		end[i] = intervals[i][1]
	}
	sort.Ints(begin)
	sort.Ints(end)
	i, j, cnt, ans := 0, 0, 0, 0
	for i < n && j < n {
		if begin[i] < end[j] {
			cnt++
			i++
		} else {
			cnt--
			j++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func main() {
	input := [][]int{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println(minMeetingRooms(input))
}
