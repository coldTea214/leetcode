func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}

	wordExist := make(map[string]bool, len(wordDict))
	sizes := make(map[int]bool)
	for _, w := range wordDict {
		sizes[len(w)] = true
		wordExist[w] = true
	}

	// dp[i] 表示 wordBreak(s[:i], wordDict)
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
	return dp[len(s)]
}
