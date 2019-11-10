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

	var doWordBreak func(int, string)
	doWordBreak = func(i int, str string) {
		if i == len(s) {
			res = append(res, str[1:])
			return
		}

		for size := range sizes {
			if i+size <= len(s) && wordExist[s[i:i+size]] {
				doWordBreak(i+size, str+" "+s[i:i+size])
			}
		}
	}
	doWordBreak(0, "")

	return res
}
