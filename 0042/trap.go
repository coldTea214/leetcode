func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	maxFromLeft := make([]int, len(height))
	maxFromLeft[0] = height[0]
	for i := 1; i < len(height); i++ {
		maxFromLeft[i] = max(height[i], maxFromLeft[i-1])
	}
	maxFromRight := make([]int, len(height))
	maxFromRight[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		maxFromRight[i] = max(height[i], maxFromRight[i+1])
	}

	var sum int
	for i := 1; i < len(height)-1; i++ {
		water := min(maxFromLeft[i-1], maxFromRight[i+1]) - height[i]
		if water > 0 {
			sum += water
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
