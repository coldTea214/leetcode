type MedianFinder struct {
	data []int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(n int) {
	i := searchInsert(mf.data, n)
	rear := append([]int{}, mf.data[i:]...)
	mf.data = append(append(mf.data[:i], n), rear...)
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func (mf *MedianFinder) FindMedian() float64 {
	if len(mf.data) == 0 {
		return 0
	}

	mid := len(mf.data) >> 1
	if len(mf.data)&1 == 1 {
		return float64(mf.data[mid])
	} else {
		return float64(mf.data[mid-1]+mf.data[mid]) / 2
	}
}
