package main

import "fmt"

func firstBadVersion(n int) int {
	low, high := 1, n
	for low < high {
		mid := (low + high) >> 1
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func isBadVersion(n int) bool {
	return n >= 4
}

func main() {
	fmt.Println(firstBadVersion(5))
}
