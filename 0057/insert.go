func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var insertLoc int
	for insertLoc = 0; insertLoc < len(intervals) && intervals[insertLoc][0] <= newInterval[0]; insertLoc++ {
	}

	var newIntervals [][]int
	newIntervals = append(newIntervals, intervals[:insertLoc]...)
	newIntervals = append(newIntervals, newInterval)
	newIntervals = append(newIntervals, intervals[insertLoc:]...)
	return merge(newIntervals)
}

func merge(intervals [][]int) [][]int {
	var res [][]int
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			end = max(end, intervals[i][1])
		} else {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	res = append(res, []int{start, end})

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
