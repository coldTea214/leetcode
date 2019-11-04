func largestRectangleArea(heights []int) int {
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	// stack 中存的是 heights 的元素的索引号
	// stack 中索引号对应的 heights 中的值，是递增的。
	// e.g.
	//     stack = []int{1,3,5}，那么
	//     heights[1] < heights[3] < heights[5]
	stack := make([]int, 1, len(heights))

	res := 0
	for end := 1; end < len(heights); {
		top := heights[stack[len(stack)-1]]
		if top < heights[end] {
			stack = append(stack, end)
			end++
			continue
		}

		begin := stack[len(stack)-2]
		index := stack[len(stack)-1]
		height := heights[index]
		/*
		   heights:
		   0  1 2 3 4 5 6 7
		   -2 2 1 5 6 2 3 -1

		   stack     begin end height
		   [0 1]     0     2   2
		   [0 2 3 4] 3     5   6
		   [0 2 3]   2     5   5
		   [0 2 5 6] 5     7   3
		   [0 2 5]   2     7   2
		   [0 2]     0     7   1
		*/
		area := (end - begin - 1) * height
		if area > res {
			res = area
		}

		// pop
		stack = stack[:len(stack)-1]
	}
	return res
}
