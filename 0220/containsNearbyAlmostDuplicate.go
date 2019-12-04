import "sort"

type Pair struct {
	val int
	idx int
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 {
		return false
	}
	pairs := make([]*Pair, len(nums))
	for i, num := range nums {
		pairs[i] = &Pair{num, i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].val != pairs[j].val {
			return pairs[i].val < pairs[j].val
		}
		return pairs[i].idx < pairs[j].idx
	})

	i, j := 0, 1
	for j < len(pairs) {
		if pairs[j].val-pairs[i].val <= t {
			if abs(pairs[j].idx-pairs[i].idx) <= k {
				return true
			}
			j++
		} else {
			i++
			if j <= i {
				j++
			}
		}
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
