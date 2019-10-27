func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	row, col := len(matrix), len(matrix[0])

	next := nextFunc(row, col)

	res := make([]int, row*col)
	for i := range res {
		x, y := next()
		res[i] = matrix[x][y]
	}

	return res
}

func nextFunc(row, col int) func() (int, int) {
	top, down := 0, row-1
	left, right := 0, col-1
	x, y := 0, -1
	dx, dy := 0, 1
	return func() (int, int) {
		x += dx
		y += dy
		switch {
		case y+dy > right:
			top++
			dx, dy = 1, 0
		case x+dx > down:
			right--
			dx, dy = 0, -1
		case y+dy < left:
			down--
			dx, dy = -1, 0
		case x+dx < top:
			left++
			dx, dy = 0, 1
		}
		return x, y
	}
}
