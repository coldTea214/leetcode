import "sort"

func threeSumSmaller(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < target {
				count, left = count+right-left, left+1
			} else {
				right--
			}
		}
	}
	return count
}
