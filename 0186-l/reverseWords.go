package main

import "fmt"

// 前置题 0151
func reverseWords(s []byte) {
	reverseWordsHelper(s, 0, len(s)-1)
	var i int
	for j := 0; j < len(s); j++ {
		if s[j] == ' ' {
			reverseWordsHelper(s, i, j-1)
			i = j + 1
		}
	}
	reverseWordsHelper(s, i, len(s)-1)
}

func reverseWordsHelper(s []byte, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	input := []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}
	reverseWords(input)
	fmt.Println(string(input))
}
