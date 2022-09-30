type Point struct {
	x int
	y int
}

func maxPoints(points [][]int) int {
	samePointCount := make(map[Point]int, len(points))

	for i := 0; i < len(points); i++ {
		point := Point{x: points[i][0], y: points[i][1]}
		samePointCount[point]++
	}

	pointNum := len(samePointCount)
	if pointNum <= 2 {
		return len(points)
	}

	// 重新生成 points，合并相同点数提高效率
	newPoints := make([]Point, 0, pointNum)
	for point := range samePointCount {
		newPoints = append(newPoints, point)
	}

	max := 0
	for i := 0; i < len(newPoints)-1; i++ {
		for j := i + 1; j < len(newPoints); j++ {
			count := 0
			for k := 0; k < pointNum; k++ {
				if isSameLine(newPoints[i], newPoints[j], newPoints[k]) {
					count += samePointCount[newPoints[k]]
				}
			}
			if max < count {
				max = count
			}
		}
	}
	return max
}

// 判断向量(p1-->p2)和向量(p1-->p3)的斜率是否相等
func isSameLine(p1, p2, p3 Point) bool {
	return (p3.x-p1.x)*(p2.y-p1.y) == (p2.x-p1.x)*(p3.y-p1.y)
}
