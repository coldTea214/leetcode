package main

import "fmt"

func quickSort(nums []int, low, high int) {
	if low < high {
		mid := partition(nums, low, high)
		quickSort(nums, low, mid-1)
		quickSort(nums, mid+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func main() {
	raw := []int{6, 3, 5, 2, 9}
	quickSort(raw, 0, 4)
	fmt.Println(raw)
}
