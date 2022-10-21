func shortestDistance(wordsDict []string, word1, word2 string) int {
	ans := len(wordsDict)
	index1, index2 := -1, -1
	for i, word := range wordsDict {
		if word == word1 {
			index1 = i
		} else if word == word2 {
			index2 = i
		}
		if index1 >= 0 && index2 >= 0 {
			ans = min(ans, abs(index1-index2))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
