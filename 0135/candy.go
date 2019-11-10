func candy(ratings []int) int {
	n := len(ratings)
	if n <= 1 {
		return n
	}

	candysFromLeft := make([]int, n)
	candysFromRight := make([]int, n)
	candysFromLeft[0] = 1
	candysFromRight[n-1] = 1

	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			candysFromLeft[i] = candysFromLeft[i-1] + 1
		} else {
			candysFromLeft[i] = 1
		}

		if ratings[n-i-1] > ratings[n-i] {
			candysFromRight[n-i-1] = candysFromRight[n-i] + 1
		} else {
			candysFromRight[n-i-1] = 1
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += max(candysFromLeft[i], candysFromRight[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
