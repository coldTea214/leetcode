package main

import "fmt"

func wordBreak(s string, wordDict []string) []string {
	wordExist := make(map[string]bool, len(wordDict))
	sizes := make(map[int]bool)
	for _, w := range wordDict {
		sizes[len(w)] = true
		wordExist[w] = true
	}

	// 不加 dp 判断，题目有个恶心的 case 会超时
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		if !dp[i] {
			continue
		}

		for size := range sizes {
			if i+size <= len(s) && wordExist[s[i:i+size]] {
				dp[i+size] = true
			}
		}
	}
	if !dp[len(s)] {
		return []string{}
	}

	var res []string
	doWordBreak(s, sizes, wordExist, "", &res)
	return res
}

func doWordBreak(s string, sizes map[int]bool, wordExist map[string]bool, str string, res *[]string) {
	if len(s) == 0 {
		*res = append(*res, str[1:])
		return
	}

	for size := range sizes {
		if size <= len(s) && wordExist[s[:size]] {
			doWordBreak(s[size:], sizes, wordExist, str+" "+s[:size], res)
		}
	}
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, wordDict))
}
