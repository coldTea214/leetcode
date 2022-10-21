import "sort"

// 一维求多点距离最多点，实际就是中位数所在点
// 二维就是两个独立的一维问题求解
func minTotalDistance(grid [][]int) int {
	var rows, cols []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	sort.Ints(rows)
	sort.Ints(cols)

	sum := 0
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		sum += rows[j] - rows[i]
		sum += cols[j] - cols[i]
	}
	return sum
}
