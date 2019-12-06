import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var res []string
	var str string
	begin := nums[0]
	for i := 0; i < len(nums); i++ {
		if i != len(nums)-1 && nums[i+1] == nums[i]+1 {
			continue
		}

		if nums[i] == begin {
			str = fmt.Sprintf("%d", begin)
		} else {
			str = fmt.Sprintf("%d->%d", begin, nums[i])
		}

		if i+1 < len(nums) {
			begin = nums[i+1]
		}
		res = append(res, str)
	}
	return res
}
