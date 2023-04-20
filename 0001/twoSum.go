// 后置题 0167
func twoSum(nums []int, target int) []int {
	numIdx := make(map[int]int)
	for idx, num := range nums {
		numIdx[num] = idx
	}

	for idx1, num := range nums {
		if idx2, ok := numIdx[target-num]; ok {
			if idx2 == idx1 {
				continue
			}
			return []int{idx1, idx2}
		}
	}
	return nil
}
