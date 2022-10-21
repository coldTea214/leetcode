import "sort"

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if end > intervals[i][0] {
			return false
		}
		end = intervals[i][1]
	}
	return true
}
