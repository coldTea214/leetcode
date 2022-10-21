package main

import (
	"fmt"
	"strings"
)

func wordPatternMatch(pattern string, s string) bool {
	res := false
	wordPatternMatchHelper(pattern, s, make(map[byte]string), make(map[string]bool), &res)
	return res
}

func wordPatternMatchHelper(pattern, s string, byte2Word map[byte]string, wordExist map[string]bool, res *bool) {
	if *res {
		return
	}
	if pattern == "" {
		if s == "" {
			*res = true
		}
		return
	}

	word, ok := byte2Word[pattern[0]]
	if !ok {
		for j := 1; j <= len(s); j++ {
			if wordExist[s[:j]] {
				continue
			}
			byte2Word[pattern[0]] = s[:j]
			wordExist[s[:j]] = true
			wordPatternMatchHelper(pattern[1:], s[j:], byte2Word, wordExist, res)
			if *res {
				return
			}
			delete(wordExist, s[:j])
			delete(byte2Word, pattern[0])
		}
	} else {
		idx := strings.Index(s, word)
		if idx != 0 {
			return
		}
		if len(word) > len(s) {
			return
		}
		wordPatternMatchHelper(pattern[1:], s[len(word):], byte2Word, wordExist, res)
	}
}

func main() {
	fmt.Println(wordPatternMatch("d", "e"))
	fmt.Println(wordPatternMatch("abab", "redblueredblue"))
	fmt.Println(wordPatternMatch("abab", "asdasdasdasd"))
}
