import "math"

type WordDistance map[string][]int

func Constructor(wordsDict []string) WordDistance {
	indicesMap := WordDistance{}
	for i, word := range wordsDict {
		indicesMap[word] = append(indicesMap[word], i)
	}
	return indicesMap
}

func (indicesMap WordDistance) Shortest(word1, word2 string) int {
	ans := math.MaxInt32
	indices1 := indicesMap[word1]
	indices2 := indicesMap[word2]
	i, n := 0, len(indices1)
	j, m := 0, len(indices2)
	for i < n && j < m {
		index1, index2 := indices1[i], indices2[j]
		ans = min(ans, abs(index1-index2))
		if index1 < index2 {
			i++
		} else {
			j++
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
