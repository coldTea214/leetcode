package main

import "fmt"

func alienOrder(words []string) string {
	neighbors := make(map[byte][]byte)
	inDegrees := make(map[byte]int)

	for _, word := range words {
		for i := range word {
			neighbors[word[i]] = []byte{}
		}
	}

	for i := 1; i < len(words); i++ {
		w1, w2 := words[i-1], words[i]
		minLen := min(len(w1), len(w2))
		j := 0
		for ; j < minLen; j++ {
			if w1[j] != w2[j] {
				neighbors[w1[j]] = append(neighbors[w1[j]], w2[j])
				inDegrees[w2[j]]++
				break
			}
		}

		// 如果没有找到边，说明较短词语是较长词语的前缀，那么正确的顺序应该是较长的词在后面，否则就是非法输入
		if j == minLen && len(w1) > len(w2) {
			return ""
		}
	}

	var order []byte
	for n := range neighbors {
		if inDegrees[n] == 0 {
			order = append(order, n)
		}
	}

	visited := 0
	for i := 0; i < len(order); i++ {
		cur := order[i]
		for _, adj := range neighbors[cur] {
			inDegrees[adj]--
			if inDegrees[adj] == 0 {
				order = append(order, adj)
			}
		}
		visited++
	}

	if visited != len(neighbors) {
		return ""
	}
	return string(order)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	input := []string{"wrt", "wrf", "er", "ett", "rftt"}
	fmt.Println(alienOrder(input))
}
