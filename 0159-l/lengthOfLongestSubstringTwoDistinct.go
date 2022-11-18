package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		oneDistinct := true
		c1, c2 := s[i], s[i]
		j := i + 1
		for ; j < len(s); j++ {
			if s[j] != c1 && s[j] != c2 {
				c2 = s[j]
				if oneDistinct {
					// 2 distinct
					oneDistinct = false
				} else {
					// 3 distinct
					break
				}
			}
		}
		if j-i > ans {
			ans = j - i
		}
		for ; i+1 < len(s); i++ {
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
