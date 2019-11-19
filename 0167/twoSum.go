func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		switch {
		case sum == target:
			return []int{low + 1, high + 1}
		case sum < target:
			low++
		case sum > target:
			high--
		}
	}
	return nil
}
