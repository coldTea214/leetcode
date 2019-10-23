func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
