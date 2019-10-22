// 执行用时 : 4 ms , 在所有 golang 提交中击败了 97.83% 的用户
// 内存消耗 : 3.8 MB , 在所有 golang 提交中击败了 29.30% 的用户

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
