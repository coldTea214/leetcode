package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		cnt := 1
		c1, c2 := s[i], s[i]
		j := i + 1
		for ; j < n; j++ {
			if s[j] != c1 && s[j] != c2 {
				c2 = s[j]
				if cnt == 1 {
					cnt--
				} else {
					break
				}
			}
		}
		if j-i > ans {
			ans = j - i
		}
		for ; i+1 < n; i++ {
			if s[i+1] != s[i] {
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
}
