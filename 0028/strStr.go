package main

import (
	"fmt"
	"strings"
)

// 解法1
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 解法2
func strStr2(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

// 解法3: kmp
// https://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
func strStr(haystack, needle string) int {
	hLen, nLen := len(haystack), len(needle)
	if nLen == 0 {
		return 0
	}
	next := make([]int, nLen)
	for i, j := 1, 0; i < nLen; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	for i, j := 0, 0; i < hLen; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == nLen {
			return i - nLen + 1
		}
	}
	return -1
}

func main() {
	// next: [0,1]
	fmt.Println(strStr("hello", "ll"))
	// next: [0,1,0]
	fmt.Println(strStr("aaaaa", "bba"))
}
