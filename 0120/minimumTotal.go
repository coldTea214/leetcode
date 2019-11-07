func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	minInRow := make([]int, n)

	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			switch {
			case j == 0:
				minInRow[0] = minInRow[0] + triangle[i][0]
			case j == i:
				minInRow[j] = minInRow[j-1] + triangle[i][j]
			case minInRow[j-1] < minInRow[j]:
				minInRow[j] = minInRow[j-1] + triangle[i][j]
			default:
				minInRow[j] = minInRow[j] + triangle[i][j]
			}
		}
	}

	min := minInRow[0]
	for i := 1; i < n; i++ {
		if min > minInRow[i] {
			min = minInRow[i]
		}
	}
	return min
}
