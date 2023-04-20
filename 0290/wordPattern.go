package main

import (
	"fmt"
	"strings"
)

// 前置题 0205
func wordPattern(pattern string, str string) bool {
	ps := strings.Split(pattern, "")
	ss := strings.Split(str, " ")
	if len(ps) != len(ss) {
		return false
	}

	pLastLoc, sLastLoc := make(map[string]int), make(map[string]int)
	for i := range ps {
		if pLastLoc[ps[i]] != sLastLoc[ss[i]] {
			return false
		}
		pLastLoc[ps[i]], sLastLoc[ss[i]] = i+1, i+1
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
}
