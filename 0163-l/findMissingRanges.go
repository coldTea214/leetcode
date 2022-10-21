package main

import (
	"fmt"
	"strconv"
)

func findMissingRanges(nums []int, lower int, upper int) (ans []string) {
	n := lower
	nums = append(nums, upper+1)
	for _, num := range nums {
		if num == n {
			n++
		} else if num > n {
			if num-1 == n {
				ans = append(ans, strconv.Itoa(n))
			} else {
				ans = append(ans, strconv.Itoa(n)+"->"+strconv.Itoa(num-1))
			}
			n = num + 1
		}
	}
	return

}

func main() {
	// [2 4->49 51->74 76->99]
	fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
}
