import (
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	numStrs := toStringArray(nums)

	quickSort(numStrs, 0, len(numStrs)-1)
	if numStrs[0] == "0" {
		return "0"
	}
	return strings.Join(numStrs, "")
}

func toStringArray(nums []int) []string {
	var strs []string
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	return strs
}

func quickSort(strs []string, low, high int) {
	if low < high {
		mid := partition(strs, low, high)
		quickSort(strs, low, mid-1)
		quickSort(strs, mid+1, high)
	}
}

func partition(strs []string, low, high int) int {
	pivot := strs[high]
	i := low - 1
	for j := low; j < high; j++ {
		// strs[j] 拼接在左边数字更大
		if strs[j]+pivot > pivot+strs[j] {
			i++
			strs[i], strs[j] = strs[j], strs[i]
		}
	}
	strs[i+1], strs[high] = strs[high], strs[i+1]
	return i + 1
}
