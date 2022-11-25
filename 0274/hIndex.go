package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)

	count := len(citations)
	low, high := 0, count-1
	var mid int
	for low <= high {
		mid = (low + high) >> 1
		if citations[mid] == count-mid {
			// 0 1 3 5 6
			return citations[mid]
		} else if citations[mid] > count-mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	// 0 1 4 5 6
	return count - low
}

func main() {
	input := []int{4, 0, 6, 1, 5}
	fmt.Println(hIndex(input))
}
