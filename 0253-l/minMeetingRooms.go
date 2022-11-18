package main

import (
	"fmt"
	"sort"
)

// 上下车问题, 等价于车上最多有多少人
// 0 -> 5 -> 10 -> 15 -> 20 -> 30
//   +1   +1    -1    +1    -1    -1
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
